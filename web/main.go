package main
import (
	"github.com/gin-gonic/gin"
	 _ "github.com/go-sql-driver/mysql"
)

func main()  {
	router := gin.Default()
	router.GET("/index", func(context *gin.Context) {

	})
}