package service

import "github.com/gogaeva/shmot-shprot/internal/repository"

// type Authorization interface {
// 	CreateUser(user domain.User) (int, error)
// 	GenerateToken(username, password string) (string, error)
// 	ParseToken(token string) (int, error)
// }

// type Look interface {
// 	MakeLook(look *domain.Look) (int, error)
// }

// type Cloth interface {
// }

type Services struct {
	Authorization *Authorization
	ClothService  *ClothService
	// Look
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthorization(repos.AuthRepo),
		ClothService:  NewClothService(repos.ClothRepo, repos.PhotoStore),
	}
}
