package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(context *gin.Context) {
		context.Header("Server", "gin")
		context.HTML(http.StatusOK, "index.html", gin.H{"title": "index"})
	})
	router.Run(":8080")
}
