package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})
	
	router.GET("/1", func(context *gin.Context) {
		context.Request.URL.Path = "/2"
		router.HandleContext(context)
	})

	router.GET("/2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	router.Run()
}
