package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com",
	})
}

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	return r
}
