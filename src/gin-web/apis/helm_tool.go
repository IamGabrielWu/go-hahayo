package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func PostMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post Message",
	})
}

// curl -X POST  -d "post_param=gabrielwu"  localhost:8080/postparam
func PostParam(c *gin.Context) {
	post_param := c.PostForm("post_param")
	fmt.Println(post_param)
	c.JSON(200, gin.H{
		"post_param": post_param,
	})
}
