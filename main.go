package main

import (
	"net/http"

	"github.com/NurymGM/New-Hotel/controllers"
	"github.com/NurymGM/New-Hotel/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.ConnectToRedis()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from root route!",
		})
	})

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.ReadPost)
	r.GET("/posts/:id", controllers.ReadPostByID)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}