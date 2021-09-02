package routers

import (
	"Gin/src/controllerUse/controllers"
	"github.com/gin-gonic/gin"
)

func AdminRouters(router *gin.Engine)  {

	admin := router.Group("/admin")
	{
		// 方法一 把匿名函数写到一个文件里变成实名的抽离出去，缺点是无法实现controller的继承
		admin.GET("/user", controllers.UserController)
	}
}
