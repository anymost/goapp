package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
	"net/http"
)

func Logger() gin.HandlerFunc  {
	return func(context *gin.Context) {
		start := time.Now()
		context.Set("name", "jack")
		context.Next()
		interval := time.Since(start)
		log.Println(interval)
		log.Println(context.Writer.Status())
	}
}

func main() {
	router := gin.Default()
	router.Use(Logger())
	router.GET("/", func(context *gin.Context) {
		name := context.MustGet("name").(string)
		context.JSON(http.StatusOK, gin.H{"name": name})
	})
	router.Run()
}
