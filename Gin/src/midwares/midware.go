package main

import (
	"Gin/src/midwares/midware"
	"Gin/src/midwares/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//配置全局的中间件，全局都会生效
	router.Use(midware.GlobleMidware{}.Show)

	routes.Midwarerouters(router)

	//goroutine 想使用中间件需要 复制一份


	router.Run(":8888")

}
