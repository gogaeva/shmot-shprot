package handler

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gogaeva/shmot-shprot/internal/domain"
)

type Look interface {
	Create(look domain.Look) (int, error)
}

type LookHandler struct {
}

func (h *Handler) createLook(c *gin.Context) {
	//var input domain.Look

	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	file, header, err := c.Request.FormFile("photo")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	description := c.Request.FormValue("description")
	season := c.Request.FormValue("season")
	purpose := c.Request.FormValue("purpose")

	input := domain.Look{
		Description: description,
		Season:      season,
		Purpose:     purpose,
	}

	c.JSON(200, &input)

	tmpfile, err := os.Create("./" + header.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer tmpfile.Close()

	_, err = io.Copy(tmpfile, file)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// if err := c.BindJSON(&input); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// id, err := h.services.MakeLook(&input)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"id": id,
	// })
}

func (h *Handler) getAllLooks(c *gin.Context) {

}

func (h *Handler) getLookById(c *gin.Context) {

}

func (h *Handler) updateLook(c *gin.Context) {

}

func (h *Handler) deleteLook(c *gin.Context) {

}

// func parseMultipart(c *gin.Context) (os.File, domain.Look, error) {

// }
