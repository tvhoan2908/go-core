package routes

import (
	"go-core/config"
	"go-core/handlers"
	"go-core/middlewares"
	"go-core/services"

	"github.com/gin-gonic/gin"
)

func (api *API) RoleRoutes(router *gin.RouterGroup) {
	handler := handlers.NewRoleHandler(services.NewRoleService(api.DB))
	router.GET("/roles", middlewares.PermissionMiddleware(config.VIEW_ALL_ROLE), handler.RoleList)
	router.POST("/roles", middlewares.PermissionMiddleware(config.ADD_NEW_ROLE), handler.CreateRole)
	router.PUT("/roles/:id", middlewares.PermissionMiddleware(config.EDIT_ANY_ROLE), handler.UpdateRole)
}
