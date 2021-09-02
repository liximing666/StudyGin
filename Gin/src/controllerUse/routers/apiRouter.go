package routers

import (
	"Gin/src/controllerUse/controllers"
	"github.com/gin-gonic/gin"
)

func ApiRouters(router *gin.Engine)  {

	api := router.Group("/api")

	{
		api.GET("/add", controllers.ApiController{}.Add)

		api.GET("/post", controllers.ApiController{}.Post)

		api.GET("/update", controllers.ApiController{}.Update)
		//继承于baseController
		api.GET("/success", controllers.ApiController{}.GO)


	}

}
