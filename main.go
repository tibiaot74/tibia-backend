package main

import (
	"fmt"

	"tibia-backend/database"
	"tibia-backend/helpers"
	"tibia-backend/routes"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title                      Tibia backend for frontend
// @version                    1.0
// @description                This is a backend API created using Go(lang) to serve a otserver Tibia website
// @host                       localhost:7474
// @BasePath                   /api
// @accept                     json
// @produce                    json
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
// @description                You can issue this token when you log in (route /api/login)
func main() {
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1),
	)
	// Get Database Env Vars
	dbUser := helpers.GetEnv("DB_USER")
	dbPassword := helpers.GetEnv("DB_PASSWORD")
	dbHost := helpers.GetEnv("DB_HOST")
	dbPort := helpers.GetEnv("DB_PORT")
	dbName := helpers.GetEnv("DB_NAME")
	// Initialize Database
	DbConnetionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	database.Connect(DbConnetionString)
	// Initialize Router
	router := routes.InitRouter()
	router.Run(":7474")
}
