package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"io"
	"os"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func bindModel() {
	gin.DisableConsoleColor()
	file, _ := os.Create("debug.log")
	gin.DefaultWriter = io.MultiWriter(file)
	router := gin.Default()

	router.POST("/loginForm", func(context *gin.Context) {
		var loginUser Login
		if err := context.ShouldBind(&loginUser); err == nil {
			fmt.Println(loginUser)
			context.JSON(http.StatusOK, loginUser)
		} else {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
		}
		context.Bind(&loginUser)
	})

	router.POST("/loginJson", func(context *gin.Context) {
		var loginUser Login
		err := context.ShouldBindJSON(&loginUser)
		fmt.Println(err)
		if err == nil {
			context.JSON(http.StatusOK, loginUser)
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
		}
	})

	router.Run()
}
