package main

import (
	"Gin/src/sessionUse/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//gin没有自带的session 需要导入第三方session
func main() {

	router := gin.Default()


	//创建session存储引擎，基于cookie
	store := cookie.NewStore([]byte("secret"))
	//配置session中间件
	router.Use(sessions.Sessions("mysession", store))

	//基于cookie的session存储
	router.GET("/setcookiesession", controller.SessionController{}.SetCookieSession)
	router.GET("/getcookiesession", controller.SessionController{}.GetCookieSession)

	//基于redis的session存储


	router.Run(":8888")


}
