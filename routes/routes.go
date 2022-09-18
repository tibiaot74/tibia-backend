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

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization", "content-type"}

	router.Use(cors.New(config))
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
			account_secured.GET("/player", controllers.ListPlayers)
			account_secured.DELETE("/player/:playerId", controllers.DeletePlayer)
		}
		base_secured := api.Group("/").Use(auth.Auth())
		{
			base_secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
