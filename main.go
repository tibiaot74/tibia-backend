package main

import (
	"fmt"
	"os"

	"tibia-backend/auth"
	"tibia-backend/controllers"
	"tibia-backend/database"

	"github.com/gin-gonic/gin"
)

func getEnv(envVarName string) string {
	if envVarName == "" {
  		return os.Getenv(envVarName)
  	}
	panic(fmt.Sprintf("The environment variable named %s must be set!", envVarName))
}

func main() {
	// Get Database Env Vars
	dbUser := getEnv("DB_USER")
	dbPassword := getEnv("DB_PASSWORD")
	dbHost := getEnv("DB_HOST")
	dbPort := getEnv("DB_PORT")
	dbName := getEnv("DB_NAME")
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
