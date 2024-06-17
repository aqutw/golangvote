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
		context.JSON(http.StatusOK, map[string]int{
			"name": 123, // note: MUST ending w ,
		})
	})

	if err := g.Run(":8080"); err != nil {
		fmt.Println("fail")
	}
}
