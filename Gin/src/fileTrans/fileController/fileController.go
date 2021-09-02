package fileController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

type FileController struct {

}

func (c FileController) ToFileLoad(context *gin.Context) {
	context.HTML(200, "fileload.html", gin.H{})
}

func (c FileController) DoFileLoad(context *gin.Context) {
	//拿到文件
	file, _ := context.FormFile("file")

	path := "./static/img/" + file.Filename
	//存文件
	context.SaveUploadedFile(file, path)

	context.JSON(200, gin.H{
		"upLoad": "ok",
		"fileName": file.Filename,
		"size": file.Size,
		"path":path,
	})

}

func (c FileController) ToMutiFileLoad(context *gin.Context) {
	context.HTML(200, "mutifileload.html", gin.H{})
}

func (c FileController) DoMutiFileLoad(context *gin.Context) {

	//拿到多个文件
	form, _ := context.MultipartForm()
	files := form.File["file"]// 注意 form.File 是 map[string][]*FileHeader 类型

	path := "./static/upload/"

	//保存多个文件
	for _, file := range files {
		context.SaveUploadedFile(file, path + file.Filename)
	}

	//按照日期分目录存文件

	context.String(200, "ok")





}

func (c FileController) ToDateMutiFileLoad(context *gin.Context) {
	context.HTML(200, "savedatefile.html", gin.H{})
}

func (c FileController) DoDateMutiFileLoad(context *gin.Context) {

	form, _ := context.MultipartForm()
	files := form.File["file"]

	path := path.Join("./static", "upload", "20210901")

	//创建目录
	os.Mkdir(path, 0666)



	for _, file := range files {

		fmt.Println(path)
		context.SaveUploadedFile(file, path + "/" +  file.Filename)
	}

	context.JSON(200, gin.H{"opt": "ok"})

}
