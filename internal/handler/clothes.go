package handler

import (
	"encoding/json"
	"image"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogaeva/shmot-shprot/internal/domain"
)

type clothService interface {
	AddCloth(cloth domain.Cloth, photo image.Image) (uint, error)
}

type ClothHandler struct {
	service clothService
}

func NewClothHandler(s clothService) *ClothHandler {
	return &ClothHandler{
		service: s,
	}
}

func (h *ClothHandler) addCloth(c *gin.Context) {
	//? maybe this block is unnecessary since
	//? Request.FormFile calls ParseMultipartForm itself
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	file, _, err := c.Request.FormFile("photo")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()
	photo, _, err := image.Decode(file)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var cloth domain.Cloth
	jsonData := c.Request.FormValue("data")
	if err := json.Unmarshal([]byte(jsonData), &cloth); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.AddCloth(cloth, photo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"cloth_id": id,
	})
}

func (h *Handler) getAllClothes(c *gin.Context) {

}

func (h *Handler) getClothById(c *gin.Context) {

}

func (h *Handler) updateCloth(c *gin.Context) {

}

func (h *Handler) deleteCloth(c *gin.Context) {

}
