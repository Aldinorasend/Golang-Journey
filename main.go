package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialiasi gin
	router := gin.Default()

	//membuat route "/" dengan method get
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Berhasil",
		})
	})
	router.Run(":3000")
}
