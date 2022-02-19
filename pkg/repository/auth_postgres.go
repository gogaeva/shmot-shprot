package repository

import (
	"fmt"

	"github.com/gogaeva/shmot-shprot/model"
	"github.com/jmoiron/sqlx"
)

// Реализует интерфейсы взаимодействия сущностей с базой

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (nickname, email, password) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Nickname, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, passwordHash string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE nickname=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, passwordHash)

	return user, err
}
