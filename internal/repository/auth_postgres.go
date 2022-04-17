package repository

import (
	"fmt"

	"github.com/gogaeva/shmot-shprot/internal/domain"
	"github.com/jmoiron/sqlx"
)

// Реализует интерфейсы взаимодействия сущностей с базой

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (uint, error) {
	var id uint
	query := fmt.Sprintf("INSERT INTO %s (nickname, email, password) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Nickname, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, passwordHash string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE nickname=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, passwordHash)

	return user, err
}
