package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginServer := gin.Default()
	//ginServer.Use(favicon.New())
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"name": "jack"})
	})
	ginServer.Run(":8082")
}
