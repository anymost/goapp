package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func mime() {
	router := gin.Default()
	router.GET("/render-xml", func(context *gin.Context) {
		name := context.DefaultQuery("name", "xml")
		context.XML(http.StatusOK, gin.H{
			"name": name,
		})
	})

	router.GET("/render-yaml", func(context *gin.Context) {
		name := context.DefaultQuery("name", "yaml")
		context.YAML(http.StatusOK, gin.H{
			"name": name,
		})
	})

	router.POST("/secure-json", func(context *gin.Context) {
		name := context.DefaultPostForm("name", "secure")
		context.SecureJSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	router.GET("/JSONP?callback=x", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		//callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	router.Run()
}
