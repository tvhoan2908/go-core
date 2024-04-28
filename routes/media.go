package routes

import (
	"go-core/handlers"
	"go-core/services"

	"github.com/gin-gonic/gin"
)

func (api *API) MediaRoutes(router *gin.RouterGroup) {
	handler := handlers.NewMediaHandler(services.NewMediaService(api.DB))
	router.POST("/media", handler.UploadFile)
	router.GET("/media", handler.MediaList)
}
