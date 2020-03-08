package main

// to import from local directory, use ". <module name>/<directory name>"
import (
	. "gin-web/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", IndexApi)
	r.POST("/", PostMessage)
	r.POST("/postparam", PostParam)
	return r
}
