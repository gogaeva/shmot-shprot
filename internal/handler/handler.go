package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gogaeva/shmot-shprot/internal/service"
)

type Handler struct {
	*AuthHandler
	*ClothHandler
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		NewAuthHandler(services.Authorization),
		NewClothHandler(services.ClothService),
	}
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

			// clothes := looks.Group("/:id/clothes")
			// {
			// 	clothes.POST("/", h.addCloth)
			// 	clothes.GET("/", h.getAllClothes)
			// 	clothes.GET("/:cloth_id", h.getClothById)
			// 	clothes.PUT("/:cloth_id", h.updateCloth)
			// 	clothes.DELETE("/:cloth_id", h.deleteCloth)
			// }
		}
		clothes := api.Group("/clothes")
		{
			clothes.POST("/", h.addCloth)
			clothes.GET("/", h.getAllClothes)
			clothes.GET("/:id", h.getClothById)
			clothes.PUT("/:id", h.updateCloth)
			clothes.DELETE("/:id", h.deleteCloth)
		}
	}
	return router
}
