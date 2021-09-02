package main

import (
	"Gin/src/fileTrans/fileController"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("./templates/**/*")

	//单个文件的上传和保存
	router.GET("/fileload", fileController.FileController{}.ToFileLoad)
	router.POST("/doloadfile", fileController.FileController{}.DoFileLoad)


	//同名的多个文件的上传和保存
	router.GET("/mutifileload", fileController.FileController{}.ToMutiFileLoad)
	router.POST("/domutiloadfile", fileController.FileController{}.DoMutiFileLoad)

	//按照日期分目录保存文件
	router.GET("/mutifileloaddate", fileController.FileController{}.ToDateMutiFileLoad)
	router.POST("/domutiloadfiledate", fileController.FileController{}.DoDateMutiFileLoad)


	router.Run(":8888")

}
