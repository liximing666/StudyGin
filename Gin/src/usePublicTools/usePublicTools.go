package main

import (
	"Gin/src/usePublicTools/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/tools", controller.ToolController{}.Todate)

	router.Run(":8888")

}
