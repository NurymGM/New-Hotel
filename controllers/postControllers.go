package controllers

import (
	"context"
	"encoding/json"
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

	// check Redis cache, if hit respond with it
	value, err := initializers.RDB.Get(context.Background(), id).Result()
	if err == nil {
		var cachedPost models.Post // Deserialize JSON string into Post struct
		err = json.Unmarshal([]byte(value), &cachedPost)
		if err == nil {
			c.IndentedJSON(http.StatusOK, cachedPost)
			return
		}
	}

	// else get post from PostgreSQL
	post := models.Post{}
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Serialize post then add it to Redis cache
	postJSON, err := json.Marshal(post)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize post"})
		return
	}

	err = initializers.RDB.Set(context.Background(), id, postJSON, 0).Err()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to set value in Redis"})
		return
	}

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
