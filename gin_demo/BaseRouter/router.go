package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
	r.POST("/xxxpost", getting)
	r.PUT("/xxxput")
	//监听端口默认为8080
	r.Run(":8000")
}
