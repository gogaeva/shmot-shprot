package service

import (
	"image"

	"github.com/gogaeva/shmot-shprot/internal/domain"
)

type ClothRepository interface {
	AddCloth(cloth domain.Cloth) (uint, error)
	DeleteCloth(id uint) error
	GetCloth(id uint) (domain.Cloth, error)
	GetAllClothes(userId uint) ([]domain.Cloth, error)
	UpdateCloth(id uint, new *domain.Cloth) error
}

type FileStore interface {
	AddPhoto(photo image.Image, userId uint) (string, error)
	DeletePhoto(userId uint, photoId string) error
}

type ClothService struct {
	repo  ClothRepository
	store FileStore
}

func NewClothService(r ClothRepository, st FileStore) *ClothService {
	return &ClothService{
		repo:  r,
		store: st,
	}
}

//Puts photo of cloth in store and makes new record in db
func (s *ClothService) AddCloth(cloth domain.Cloth, photo image.Image) (uint, error) {
	photoId, err := s.store.AddPhoto(photo, cloth.OwnerId)
	if err != nil {
		return 0, err
	}

	cloth.PhotoId = photoId

	clothId, err := s.repo.AddCloth(cloth)
	if err != nil {
		return 0, err
	}

	return clothId, nil
}
