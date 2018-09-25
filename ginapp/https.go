package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/autotls"
	"net/http"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"name": "jack",
		})
	})
	log.Fatal(autotls.Run(router, "https.go.com"))
}
