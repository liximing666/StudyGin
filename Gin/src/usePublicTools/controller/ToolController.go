package controller

import (
	"Gin/tools"
	"github.com/gin-gonic/gin"
)

type ToolController struct {


}

func (c ToolController) Todate(context *gin.Context) {

	//调用全局的工具
	S := tools.ToDate(1313136988)

	context.String(200, S)
}
