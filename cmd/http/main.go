package main

import (
	"go-core/config"
	Middlewares "go-core/middlewares"
	"go-core/routes"
	Utils "go-core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	_ "go-core/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Core API Docs
// @version         1.0
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(Middlewares.ErrorHandler())

	router.NoRoute(func(c *gin.Context) {
		c.Error(Middlewares.NewError(http.StatusNotFound, "Page not found."))
	})

	// Config upload file
	// 8 << 20 = 2^20 * 8
	// 1 << 3 = 2^3 * 1
	router.MaxMultipartMemory = 8 << 20
	router.Static("/public", "./public")

	config.ConnectRedis()
	db := config.ConnectDatabase()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unique_name", Utils.UniqueName)
	}

	// Cors
	router.Use(Middlewares.CorsMiddleware())
	api := &routes.API{
		DB:     db,
		Router: router,
	}
	api.SetupRouter()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go Utils.SendMails()

	router.Run()
}
