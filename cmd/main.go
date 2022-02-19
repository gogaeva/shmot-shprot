package main

import (
	"log"
	"os"

	model "github.com/gogaeva/shmot-shprot"
	"github.com/gogaeva/shmot-shprot/pkg/handler"
	"github.com/gogaeva/shmot-shprot/pkg/repository"
	"github.com/gogaeva/shmot-shprot/pkg/service"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	//_ "github.com/jackc/pgx"
	"github.com/spf13/viper"
)

func main() {
	// handlers := new(handler.Handler)
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config, %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRpository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(model.Server)
	port := viper.GetString("port")
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
