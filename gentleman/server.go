package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()
	router.GET("/json", func(context *gin.Context) {
		fmt.Println(context.GetHeader("client"))
		context.Header("server", "gin")
		context.SetCookie("id", "abcdefg",10000, "/", "localhost", false, false)
		context.JSON(http.StatusOK, gin.H{
			"name": "jack",
		})
	})
	router.Run(":5000")
}
