package repository

import (
	"fmt"

	//_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

// Создает объект базы и устанавливает соединение

const (
	usersTable        = "users"
	clothesTable      = "clothes"
	looksTable        = "looks"
	looksClothesTable = "looks_clothes"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode,
	))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
