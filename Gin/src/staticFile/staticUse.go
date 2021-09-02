package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/**/*")//配置导入模板所在的目录
	engine.Static("/static", "./static")//配置静态资源的路径

	//注册模板函数

	
	engine.GET("/statictest", func(context *gin.Context) {
		context.HTML(200, "static.html", map[string]interface{}{})
	})


	engine.Run(":6666")
	

}
