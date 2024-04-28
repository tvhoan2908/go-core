package routes

import (
	"go-core/config"
	"go-core/handlers"
	"go-core/middlewares"
	"go-core/services"

	"github.com/gin-gonic/gin"
)

func (api *API) CategoryRoutes(router *gin.RouterGroup) {
	handler := handlers.NewCategoryHandler(services.NewCategoryService(api.DB))
	router.POST("/categories", middlewares.PermissionMiddleware(config.ADD_NEW_CATEGORY), handler.CreateCategory)
	router.PUT("/categories/:id", middlewares.PermissionMiddleware(config.EDIT_ANY_CATEGORY), handler.UpdateCategory)
	router.GET("/categories", middlewares.PermissionMiddleware(config.VIEW_ALL_CATEGORY), handler.CategoryList)
}
