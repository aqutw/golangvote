package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/login", func(context *gin.Context) {
		fmt.Println("hihi")
	})

	if err := g.Run(":8080"); err != nil {
		fmt.Println("fail")
	}
}
