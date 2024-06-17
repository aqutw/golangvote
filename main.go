package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", getDBPwd(), "127.0.0.1:3306", "golangvote")
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:", err)
		panic(err)
	}

	//
	g := gin.Default()
	g.LoadHTMLGlob("tmpl/*")
	g.GET("/login", func(context *gin.Context) {
		fmt.Println("hihi")
		context.HTML(http.StatusOK, "login.tmpl", nil)
	})
	g.POST("/login", func(context *gin.Context) {
		var u User
		//ret := make(map[string]any)
		var ret User
		_ = context.ShouldBind(&u)
		conn.Debug()
		conn.Table("user").Where("name=?", u.Name).First(&ret)
		err := conn.Error
		if err != nil {
			fmt.Println("err", err)
			context.JSON(http.StatusBadGateway, map[string]string{
				"message": "database query fail",
			})
			return
		}
		context.JSON(http.StatusOK, ret)
	})

	if err := g.Run(":8080"); err != nil {
		fmt.Println("fail")
	}
}

type User struct {
	//	gorm.Model
	Id           int       `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" form:"name"`
	Password     string    `json:"password" form:"password"`
	CreatedTime  time.Time `json:"created_time" gorm:""`
	UpdateedTime time.Time `json:"updated_time" gorm:""`
}

func getDBPwd() string {
	pwd, err := os.ReadFile("./pwd.txt")
	if err != nil {
		fmt.Println("getDBPwd err", err)
		os.Exit(0)
	}
	strpwd := string(pwd)
	strpwd = strings.TrimSuffix(strpwd, "\n")
	strpwd = strings.TrimSuffix(strpwd, "\r")
	return strpwd
}
