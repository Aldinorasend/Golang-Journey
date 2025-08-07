package main

import (
	"golang_be/starting/controllers"
	"golang_be/starting/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialiasi gin
	router := gin.Default()
	models.ConnectDB()
	//membuat route "/" dengan method get
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Berhasil",
		})
	})
	//definisi port
	router.GET("/api/posts", controllers.FindPosts)
	router.Run(":3000")
}
