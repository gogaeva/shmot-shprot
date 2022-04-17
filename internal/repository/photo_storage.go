package repository

import (
	"image"
	"image/jpeg"
	"os"
	"path"
	"strconv"

	"github.com/segmentio/ksuid"
)

//For now storage is just a directory in local fylesystem
type FileStore struct {
	path string
}

func NewFileStore(path string) *FileStore {
	return &FileStore{path: path}
}

//Generates id for photo and save it in .jpeg file
func (st *FileStore) AddPhoto(photo image.Image, userId uint) (string, error) {
	userDir := path.Join(st.path, strconv.FormatUint(uint64(userId), 10))
	photoId := ksuid.New().String()
	fileName := photoId + ".jpeg"

	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		return "", err
	}

	filePath := path.Join(userDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	jpeg.Encode(file, photo, nil)

	return photoId, nil
}

func (st *FileStore) DeletePhoto(userId uint, photoId string) error {
	userDir := path.Join(st.path, strconv.FormatUint(uint64(userId), 10))
	fileName := photoId + ".jpeg"
	filePath := path.Join(userDir, fileName)

	return os.Remove(filePath)
}
