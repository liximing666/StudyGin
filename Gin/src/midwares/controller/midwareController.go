package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type MidwareController struct {

}

func (midc MidwareController) Show(context *gin.Context)  {
	context.String(200, "config ok")
}

func (midc MidwareController) End(context *gin.Context)  {
	context.String(200, "midware run end")
}

func (midc MidwareController) ShowData(context *gin.Context) {
	ip, _ := context.Get("ip")

	fmt.Printf("ip : %s\n", ip)
}