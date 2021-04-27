<center><h1>apply-work</h1></center>

- 记录、沉淀

[TOC]

https://github.com/nateshao/apply-work

或者可以访问这个国内gitee：https://gitee.com/nateshao/apply-work

# 第一周

## day01--2021-04-22 作业1

> 总结：配置环境。输出hello，同时在Linux服务器上，配置go环境，同时输出hello world

![001_hello.go](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210422231117678.png)

![001_hello-linux.go](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210423001738446.png)

## day02--2021-04-23 作业2

![image-20210425211136868](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210425211136868.png)

![image-20210424124655156](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210424124655156.png)

## day04--2021-04-25  作业3

![image-20210425205508655](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210425205508655.png)



![image-20210425235557611](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210425235557611.png)

# 搭建Post、Get

## get请求

![image-20210426095431167](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210426095431167.png)

## post请求

![image-20210426095527484](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210426095527484.png)

post请求原因：Apache、IIS、Nginx等绝大多数web服务器，都不允许静态文件响应POST请求，否则会返回“HTTP/1.1 405 Method not allowed”错误。

## 数据库设计

```go
CREATE TABLE `liuyan` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uname` CHAR(20) DEFAULT NULL, 			// uname 用户名
  `email` VARCHAR(32) NOT NULL,				// 邮箱
  `content` TEXT,  				 			// content 留言内容
  `insert_time` INT(10) UNSIGNED NOT NULL,	// insert_time 留言时间
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8

CREATE TABLE `user` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uname` CHAR(20) DEFAULT NULL, 			
  PRIMARY KEY (`id`)
  foreign key(liuyan_id) references dep(id)
) ENGINE=INNODB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8
```

## main.go

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type User struct {
	Id   int    `gorm:"AUTO_INCREMENT"` // 自增
	Name string `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设
	Age  int
}

var (
	db  *gorm.DB
	err error
)

func main() {
	// 链接 postgres
	db, err = gorm.Open("postgres", "port=5432 user=postgres password=123456 dbname=liuyan sslmode=disable")
	//还可以通过这种方式打开
	//db, err := sql.Open("postgres", "postgres://pqgotest:123456@localhost/liuyan?sslmode=verify-full")

	// 链接 mysql
	//db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	} else {
		db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

		if !db.HasTable(&User{}) {
			if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}).Error; err != nil {
				panic(err)
			}
		}
	}

	// 引入路由
	Router()
}

func Router() {
	router := gin.Default()
	// 路径映射
	router.GET("/user", InitPage)
	router.GET("/user/list", ListUser)
	router.GET("/user/find/:id", GetUser)
	router.POST("/user/create", CreateUser)
	router.POST("/user/upload",UploadUser)
	router.PUT("/user/update/:id", UpdateUser)
	router.DELETE("/user/:id", DeleteUser)
	router.Run(":8080")
}

func InitPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK!",
	})
}


// 比如用户留言附加截图等
func UploadUser(c *gin.Context)  {
	r := gin.Default()
	// 对上传的文件做出限制
	r.MaxMultipartMemory = 8 << 20

	file, err := c.FormFile("file")
	if err != nil {
		c.String(500,"上传图片出错啦。")
	}
	c.SaveUploadedFile(file,file.Filename)
	c.String(http.StatusOK,file.Filename)
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user) // 使用bindJson填充数据

	// db.Create(&user)	// 创建对象
	// c.JSON(http.StatusOK, &user)	// 返回页面
	if user.Name != "" && user.Age > 0 {
		db.Create(&user)
		c.JSON(http.StatusOK, gin.H{"success": &user})

	}
	return

	//else {
	//	c.JSON(422, gin.H{"error": "Fields are empty"})
	//}
	//
}

// 更新用户
// http://localhost:8080/user/update/5
func UpdateUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	err := db.First(&user, id).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	} else {
		c.BindJSON(&user)
		db.Save(&user)               // 提交修改
		c.JSON(http.StatusOK, &user) // 返回页面
	}
}

// 列出所有用户
// http://127.0.0.1:8080/user/list
// curl -i http://localhost:8080/user/list
func ListUser(c *gin.Context) {
	var user []User
	db.Find(&user)
	c.JSON(http.StatusOK, &user)
}

// 列出单个用户
// curl -i http://localhost:8080/user/find/4
func GetUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	err := db.First(&user, id).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, &user)
	}
}

// 删除用户
// curl -i -X DELETE http://localhost:8080/user/1
func DeleteUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	db.First(&user, id)
	if user.Id != 0 {
		db.Delete(&user)
		c.JSON(http.StatusOK, gin.H{
			"success": "User# " + id + " deleted!",
		})
	} else {
		c.JSON(404, gin.H{
			"error": "User not found",
		})
	}
}
```

- 启动成功

![image-20210426160310063](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210426160310063.png)

- 运用Golang ide ![image-20210426163551860](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210426163551860.png)

进行发送get和post请求

## post

![image-20210426162635737](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210426162635737.png)

## Get

![image-20210426162748723](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210426162748723.png)





## MySQL换为PostgreSQL

![image-20210427092324880](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/image-20210427092324880.png)