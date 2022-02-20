package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gogaeva/shmot-shprot/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		looks := api.Group("/looks")
		{
			looks.POST("/", h.createLook)
			looks.GET("/", h.getAllLooks)
			looks.GET("/:id", h.getLookById)
			looks.PUT("/:id", h.updateLook)
			looks.DELETE("/:id", h.deleteLook)

			clothes := looks.Group("/:id/clothes")
			{
				clothes.POST("/", h.addCloth)
				clothes.GET("/", h.getAllClothes)
				clothes.GET("/:cloth_id", h.getClothById)
				clothes.PUT("/:cloth_id", h.updateCloth)
				clothes.DELETE("/:cloth_id", h.deleteCloth)
			}
		}
	}
	return router
}
