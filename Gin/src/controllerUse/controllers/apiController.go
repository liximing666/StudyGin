package controllers

import (
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	  BaseController
}

//把路由中处理业务的个个方法绑定到结构体，这样可以继承
func (apic ApiController) Add(context *gin.Context)  {
	context.String(200, "add")
}

func (apic ApiController) Post(context *gin.Context)  {
	context.String(200, "post")
}

func (apic ApiController) Update(context *gin.Context)  {
	context.String(200, "update")
}

func (apic ApiController) GO(context *gin.Context)  {
	apic.Success(context)
}
