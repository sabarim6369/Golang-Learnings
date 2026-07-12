// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	// Create Gin router
// 	router := gin.Default()

// 	// Create a GET API
// 	router.GET("/hello", func(c *gin.Context) {

// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Hello Gin",
// 		})

// 	})

// 	// Start server on port 8080
// 	router.Run(":8080")
// }

// package main
// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )
// func main(){
// 	router:=gin.Default()
// 	router.GET("/hello",func(c *gin.Context){
// 		c.JSON(http.StatusOK,gin.H{
// 			"message":"Hello Gin",
// 		})
// 	})
// 	router.POST("/login",func(c *gin.Context){
// 		c.JSON(http.StatusOK,gin.H{
// 			"message":"Login Successful",
// 		})
// 	})
// 	router.GET("/user/:name",func(c *gin.Context){
// 		name:=c.Param("name")
// 		c.JSON(http.StatusOK,gin.H{
// 			"message":"Hello " + name,
// 		})
// 	})
// 	router.Run(":8080")
// }
package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	ConnectDB()

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {

		collection := DB.Collection("students")

		student := bson.M{
			"name":  "Sabari",
			"age":   21,
			"course": "CSE",
		}

		result, err := collection.InsertOne(context.Background(), student)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Student Added Successfully",
			"id":      result.InsertedID,
		})
	})

	router.Run(":8080")
}