package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.LoadHTMLGlob("tmpl/*")
	g.GET("/login", func(context *gin.Context) {
		fmt.Println("hihi")
		context.HTML(http.StatusOK, "login.tmpl", nil)
	})
	g.POST("/login", func(context *gin.Context) {
		var u User
		_ = context.ShouldBind(&u)
		context.JSON(http.StatusOK, u)
	})

	if err := g.Run(":8080"); err != nil {
		fmt.Println("fail")
	}
}

type User struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}
