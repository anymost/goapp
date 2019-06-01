package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main()  {
	router := gin.Default()

	router.StaticFile("/",)

	router.GET("/login", func(context *gin.Context) {
		var user *User
		err := context.ShouldBindJSON(user)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		} else {
			fmt.Println(user)
		}
	})
}
