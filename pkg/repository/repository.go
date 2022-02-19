package repository

import (
	"github.com/gogaeva/shmot-shprot/model"
	"github.com/jmoiron/sqlx"
)

// Объявляет интерфейсы взаимодействия сущностей с репозиторием
// Объявляет структуру Repository, которая агрегирует реализации этих интерфейсов

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Look interface {
}

type Cloth interface {
}

type Repository struct {
	Authorization
	Look
	Cloth
}

func NewRpository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
