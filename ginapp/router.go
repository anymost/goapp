package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	router := gin.Default()
	router.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"name": "jack"})
	})

	router.GET("/string", func(context *gin.Context) {
		context.String(http.StatusOK, "hello string")
	})

	router.GET("/user/:id", func(context *gin.Context) {
		userId := context.Param("id")
		userName := context.Query("firstName")
		context.String(http.StatusOK, "hello" + userId + " hi " + userName)
	})

	router.POST("/submit", func(context *gin.Context) {
		name := context.DefaultPostForm("name", "jack")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{"name": name, "password": password})
	})

	router.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		fmt.Println(file)
		context.JSON(http.StatusOK, gin.H{"code": "ok"})

	})

	v1 := router.Group("/v1")
	{
		v1.GET("/a", func(context *gin.Context) {
			path := context.GetHeader("Accept")
			context.JSON(http.StatusOK, gin.H{"path": path})
		})
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/b", func(context *gin.Context) {
			context.Header("Access-Control-Allow-Origin", "*")
			context.Header("Location", "http://www.baidu.com")

			context.String(http.StatusMovedPermanently, "move")
		})
	}

	router.Run()
}
