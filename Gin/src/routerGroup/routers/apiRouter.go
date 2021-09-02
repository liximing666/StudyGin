package routers

import "github.com/gin-gonic/gin"

func ApiRouterInit(router *gin.Engine)  {

	api := router.Group("/api")
	api.GET("/add", func(context *gin.Context) {
		context.String(200, "add")
	})

}