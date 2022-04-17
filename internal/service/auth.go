package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogaeva/shmot-shprot/internal/domain"
)

const (
	salt       = "hhsdahuhku"
	signingKey = "ddaoshoijfwqaj"
	tokenTTL   = time.Hour * 12
)

type authRepository interface {
	CreateUser(user domain.User) (uint, error)
	GetUser(username, password string) (domain.User, error)
}

type Authorization struct {
	repo authRepository
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func NewAuthorization(r authRepository) *Authorization {
	return &Authorization{repo: r}
}

func (s *Authorization) CreateUser(user domain.User) (uint, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *Authorization) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Authorization) ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
