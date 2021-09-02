package controllers

import "github.com/gin-gonic/gin"

type BaseController struct {

}

func (baseC BaseController) Success(context *gin.Context)  {
	context.String(200, "success")
}
