package routers

import "github.com/gin-gonic/gin"

//入参是路由 engine
func DefaultRouterInit(router *gin.Engine)  {
	defaultRouter := router.Group("/")
	{
		defaultRouter.GET("/", func(context *gin.Context) {
			context.String(200, "index")
		})

		defaultRouter.GET("/news", func(context *gin.Context) {
			context.String(200, "news")
		})
	}
	
}