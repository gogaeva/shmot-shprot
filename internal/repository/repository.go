package repository

import (
	"github.com/jmoiron/sqlx"
)

// Объявляет интерфейсы взаимодействия сущностей с репозиторием
// Объявляет структуру Repository, которая агрегирует реализации этих интерфейсов

// type Authorization interface {
// 	CreateUser(user domain.User) (int, error)
// 	GetUser(username, password string) (domain.User, error)
// }

// type Look interface {
// 	CreateLook(look domain.Look) (int, error)
// }

// type Cloth interface {
// }

type Repository struct {
	AuthRepo *AuthPostgres
	//LookRepo *LookPostgres
	ClothRepo  *ClothRepo
	PhotoStore *FileStore
}

func NewRepository(db *sqlx.DB, store *FileStore) *Repository {
	return &Repository{
		AuthRepo:   NewAuthPostgres(db),
		ClothRepo:  NewClothRepo(db),
		PhotoStore: store,
		//LookRepo: NewLookPostgres(db),
	}
}
