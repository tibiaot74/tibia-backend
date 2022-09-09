package main

import (
	"fmt"

	"tibia-backend/auth"
	"tibia-backend/controllers"
	"tibia-backend/database"
	"tibia-backend/helpers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get Database Env Vars
	dbUser := helpers.GetEnv("DB_USER")
	dbPassword := helpers.GetEnv("DB_PASSWORD")
	dbHost := helpers.GetEnv("DB_HOST")
	dbPort := helpers.GetEnv("DB_PORT")
	dbName := helpers.GetEnv("DB_NAME")
	// Initialize Database
	db_connection_string := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	database.Connect(db_connection_string)
	// Initialize Router
	router := initRouter()
	router.Run(":7474")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/account", controllers.RegisterAccount)
		api.GET("/health", controllers.HealthCheck)
		secured := api.Group("/secured").Use(auth.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
