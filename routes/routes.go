package routes

import (
	"tibia-backend/auth"
	"tibia-backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")
	{
		api.POST("/login", controllers.GenerateToken)
		api.POST("/account", controllers.RegisterAccount)
		api.GET("/health", controllers.HealthCheck)
		secured := api.Group("/secured").Use(auth.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			secured.POST("/player", controllers.RegisterPlayer)
		}
	}
	return router
}
