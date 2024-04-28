package routes

import (
	"go-core/config"
	"go-core/handlers"
	"go-core/middlewares"
	"go-core/services"

	"github.com/gin-gonic/gin"
)

func (api *API) UserRoutes(router *gin.RouterGroup) {
	handler := handlers.NewUserHandler(services.NewUserService(api.DB))
	router.GET("/user", handlers.UserInfo)
	router.POST("/users", middlewares.PermissionMiddleware(config.ADD_NEW_USER), handler.StoreUser)
	router.GET("/users", middlewares.PermissionMiddleware(config.VIEW_ALL_USER), handler.UserList)
	router.PUT("/users/:id", middlewares.PermissionMiddleware(config.EDIT_ANY_USER), handler.UpdateUser)
	router.PATCH("/user/change-password", handler.ChangePassword)
}
