package main

import (
	"book/db"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //init()
	"net/http"
)

func main() {

	err := db.InitDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
	}
	fmt.Println("连接数据库成功！")

	r := gin.Default()
	r.LoadHTMLGlob("../templates/*")
	r.GET("/book_list", func(c *gin.Context) {
		dataAll, err := db.QueryAllBook()
		if err != nil {
			fmt.Printf("get data failed:err", err)
		}
		c.HTML(http.StatusOK, "book_list.html", gin.H{"data": dataAll})
	})
	//添加新书
	r.GET("/book/new", func(c *gin.Context) {
		c.HTML(200, "new_book.html", nil)
	})
	r.POST("/book/new", func(c *gin.Context) {
		tile := c.PostForm("title")
		price := c.PostForm("price")
		db.AddBook(tile, price)
		//c.String(http.StatusOK,"add book successful!")
		c.Redirect(http.StatusMovedPermanently, "/book_list")
	})

	//删除书箱
	r.GET("/book/delete", func(c *gin.Context) {
		id := c.Query("id")
		db.DelBook(id)
		c.String(http.StatusOK, "delete book successful!")
	})
	r.Run(":8999")

}
