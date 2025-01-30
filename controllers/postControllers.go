package controllers

import (
	"net/http"

	"github.com/NurymGM/New-Hotel/initializers"
	"github.com/NurymGM/New-Hotel/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// get data off req body
	body := models.Post{}
	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Couldnt create a post"})
        return
	}

	// return it
	c.IndentedJSON(http.StatusCreated, post)
}

func ReadPost(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond with them
	c.IndentedJSON(http.StatusOK, posts)
}

func ReadPostByID(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get the post
	post := models.Post{}
	initializers.DB.First(&post, id)

	// respond with it
	c.IndentedJSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data off req body
	body := models.Post{}
	c.Bind(&body)

	// find the post we are updating
	post := models.Post{}
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// respond with it
	c.IndentedJSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}