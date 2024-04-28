package routes

import (
	"go-core/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type API struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (api *API) SetupRouter() {
	v1 := api.Router.Group("/api/v1")
	api.AuthRouter(v1)

	v1Admin := api.Router.Group("/api/v1/admin")

	v1Admin.Use(middlewares.AuthMiddleware())
	{
		api.CategoryRoutes(v1Admin)
		api.RoleRoutes(v1Admin)
		api.MediaRoutes(v1Admin)
		api.UserRoutes(v1Admin)
	}
}
