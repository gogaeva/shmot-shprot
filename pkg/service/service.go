package service

import (
	"github.com/gogaeva/shmot-shprot/model"
	"github.com/gogaeva/shmot-shprot/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Look interface {
}

type Cloth interface {
}

type Service struct {
	Authorization
	Look
	Cloth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
