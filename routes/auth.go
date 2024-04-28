package routes

import (
	"go-core/handlers"
	"go-core/services"

	"github.com/gin-gonic/gin"
)

func (api *API) AuthRouter(router *gin.RouterGroup) {
	handler := handlers.NewAuthHandler(services.NewAuthService(api.DB))
	router.POST("/auth/register", handler.Register)
	router.POST("/auth/login", handler.Login)
}
