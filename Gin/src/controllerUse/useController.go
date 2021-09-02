package main

import (
	"Gin/src/controllerUse/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()


	//方法一 把匿名函数写到一个文件里变成实名的抽离出去，缺点是无法实现controller的继承
	routers.AdminRouters(router)

	//方法二 把路由中处理业务的个个方法绑定到结构体，这样可以继承，推荐使用！
	routers.ApiRouters(router)





	router.Run(":8888")

}
