package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// gin based on net/http and each request handled by individual goroutine
	router := gin.Default()
	router.GET("/async/:id", func(context *gin.Context) {
		contextCopy := context.Copy()
		go func() {
			id := contextCopy.Param("id")
			fmt.Println(id)
			context.JSON(http.StatusOK, gin.H{"id": id})
		}()
	})

	router.GET("/sync/:id", func(context *gin.Context) {
		id := context.Param("id")
		fmt.Println(id)
		time.Sleep(5 * time.Second)
		context.JSON(http.StatusOK, gin.H{"id": id})
	})

	router.Run()
}
