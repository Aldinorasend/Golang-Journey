package controllers

import (
	"golang_be/starting/models"

	"github.com/gin-gonic/gin"
)

func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(200, gin.H{
		"success": true,
		"message": "list data posts",
		"data":    posts,
	})
}
