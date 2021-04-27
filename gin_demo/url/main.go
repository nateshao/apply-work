package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "千羽")
		c.String(http.StatusOK, fmt.Sprintf("hello qianyu", name))
	})
	r.Run()
}
