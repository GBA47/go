package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	//当一个HTTP请求，用GET方法访问的时侯，如果访问路径是 /hello，
	server.GET("/hello", func(c *gin.Context) {
		// 就执行这段代码
		c.String(http.StatusOK, "hello go")
	})
	server.POST("/Post", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, post方法")
	})
	server.GET("/user/:name", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, 这是参数路由")
	})
	server.GET("/views/*.html", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, 这是通配符路由")
	})
	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

}
