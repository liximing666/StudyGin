package main

import (
	"Gin/src/routerGroup/routers"
	"github.com/gin-gonic/gin"
)

//路由的分组管理
func main() {
	router := gin.Default()

	//配置分组路由的文件
	routers.AdminRouterInit(router)
	routers.DefaultRouterInit(router)
	routers.ApiRouterInit(router)



	router.Run(":8888")
}
