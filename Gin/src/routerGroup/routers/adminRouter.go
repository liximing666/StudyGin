package routers

import "github.com/gin-gonic/gin"

func AdminRouterInit(router *gin.Engine)  {

	admin := router.Group("/admin")
	{
		admin.GET("/user", func(context *gin.Context) {
			context.String(200, "user")
		})
	}
}
