package routes

import (
	"tibia-backend/auth"
	"tibia-backend/controllers"

	docs "tibia-backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")

	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	{
		api.POST("/login", controllers.GenerateToken)
		api.POST("/account", controllers.RegisterAccount)
		api.GET("/health", controllers.HealthCheck)
		account_secured := api.Group("/account").Use(auth.Auth())
		{
			account_secured.POST("/player", controllers.RegisterPlayer)
		}
		api.GET("/ping", controllers.Ping).Use(auth.Auth())
	}
	return router
}
