package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "您已成功访问到TestServer的1122 端口")
	})
	//3.监听端口，默认8080
	r.Run(":1122")
}
