package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct{}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	server.POST("/users/signup", u.SignUp)
	server.POST("/users/login", u.Login)
	server.POST("/users/edit", u.Edit)
	server.GET("/users/profile", u.Profile)
}

type SignUpReq struct {
	Email           string `json:"email"`
	ConfirmPassword string `json:"confirmPassword"`
	Password        string `json:"password"`
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	var req SignUpReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 这里你可以添加进一步的处理逻辑，例如检查密码是否匹配
	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password and Confirm Password do not match"})
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	fmt.Printf("SignUp Request: %+v\n", req)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	// 登录逻辑
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	// 编辑用户信息逻辑
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	// 获取用户资料逻辑
}

func main() {
	r := gin.Default()

	handler := &UserHandler{}
	handler.RegisterRoutes(r)

	r.Run(":8080")
}
