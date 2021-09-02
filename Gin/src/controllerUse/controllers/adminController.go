package controllers

import "github.com/gin-gonic/gin"

//抽离出的匿名函数
func UserController(context *gin.Context) {
	context.String(200, "user")
}