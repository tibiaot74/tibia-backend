package main

import (
	"fmt"

	"tibia-backend/database"
	"tibia-backend/helpers"
	"tibia-backend/routes"
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
	router := routes.InitRouter()
	router.Run(":7474")
}
