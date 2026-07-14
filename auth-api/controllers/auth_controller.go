package controllers

import (
	"auth-api/models"
	"auth-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	err := services.Signup(user)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Signup Successful",
	})
}

func Login(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	err := services.Login(user.Email, user.Password)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
	})
}

func ChangePassword(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Change Password API",
	})
}

func Profile(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile API",
	})
}

func Logout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Successful",
	})
}