package main

import (
	"auth-api/config"
	"auth-api/database"
	"auth-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()

	database.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}