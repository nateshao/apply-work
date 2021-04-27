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
	router.POST("/login", LoginUser)
	router.POST("/register", register)
	router.POST("/user/create", CreateUser)
	router.POST("/user/upload", UploadUser)
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
func UploadUser(c *gin.Context) {
	r := gin.Default()
	// 对上传的文件做出限制
	r.MaxMultipartMemory = 8 << 20

	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错啦。")
	}
	c.SaveUploadedFile(file, file.Filename)
	c.String(http.StatusOK, file.Filename)
}

// 用户登录
func LoginUser(c *gin.Context) {
	c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, type))
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
