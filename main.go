package main

import (
	"tibia-backend/auth"
	"tibia-backend/controllers"
	"tibia-backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("root:YES@tcp(localhost:3306)/tibia?parseTime=true")
	// Initialize Router
	router := initRouter()
	router.Run(":7474")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/account", controllers.RegisterAccount)
		secured := api.Group("/secured").Use(auth.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
