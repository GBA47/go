package main

import (
	"github.com/gin-gonic/gin"
	"learn/basic-go/webook/internal/web"
)

func main() {
	server := gin.Default()
	u := &web.UserHandler{}

	server.POST("/users/signup", u.SignUp)
	//REST风格
	//server.POST("/users", func(Context *gin.Context) {
	//
	//})

	server.POST("/users/login", u.Login)
	server.POST("/users/Edit", u.Edit)
	//REST风格
	//server.POST("/users/：id", func(Context *gin.Context) {
	//
	//})

	server.GET("/users/profile", u.Profile)

	server.Run(":8080")
}
