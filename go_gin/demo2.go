package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() // 返回一个默认的gin实例
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run() // 默认在 0.0.0.0:8080 上监听并服务
}
