package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogaeva/shmot-shprot/internal/domain"
)

// Принимает данные с клиента, отправляет их в слой бизнес логики, возвращает ответ
type authorization interface {
	CreateUser(user domain.User) (uint, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (uint, error)
}

type AuthHandler struct {
	service authorization
}

func NewAuthHandler(s authorization) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) signUp(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(input.Nickname, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
