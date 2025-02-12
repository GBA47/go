package main

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()
	gob.RegisterRoutes(server)
	return server
}
func registerUsersRoutes(server *gin.Engine) {
	u := &UserHandler{}
	server.POST("/users/signup", u.Signup)
	server.POST("/users/login", u.Login)
	server.POST("users/edit", u.Edit)
	server.GET("users/profile", u.profile)
}
