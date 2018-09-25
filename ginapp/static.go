package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
)

func main() {

	router := gin.Default()
	router.Static("/static", "./assets")

	router.GET("/api/image", func(context *gin.Context) {
		resp, err := http.Get("https://ss1.baidu.com/6ONXsjip0QIZ8tyhnq/it/u=3811938961,201576138&fm=173&app=25&f=JPEG?w=218&h=146&s=2280D44D72781A3E0EA4ADBA03001003")
		if err != nil || resp.StatusCode != http.StatusOK {
			fmt.Println(err)
			return
		}
		render := resp.Body
		contentType := resp.Header.Get("Content-Type")
		contentLength := resp.ContentLength
		//headers := map[string]string{
		//	"Content-Disposition": `attachment; filename="image.png"`,
		//}
		headers := map[string]string {
			"StatusCode": strconv.Itoa(http.StatusOK),
		}
		context.DataFromReader(http.StatusOK, contentLength, contentType, render, headers)
	})
	router.Run()
}
