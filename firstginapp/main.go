package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create Gin router
	router := gin.Default()

	// Create a GET API
	router.GET("/hello", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Gin",
		})

	})

	// Start server on port 8080
	router.Run(":8080")
}