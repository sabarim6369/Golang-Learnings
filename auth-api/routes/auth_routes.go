package routes

import (
	"auth-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/signup", controllers.Signup)

	router.POST("/login", controllers.Login)

	router.PUT("/change-password", controllers.ChangePassword)

	router.GET("/profile", controllers.Profile)

	router.POST("/logout", controllers.Logout)
}