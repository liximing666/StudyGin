package routes

import (
	"Gin/src/midwares/controller"
	midware2 "Gin/src/midwares/midware"
	"github.com/gin-gonic/gin"
)

func Midwarerouters(router *gin.Engine)  {

	midware := router.Group("/midware", midware2.Midware{}.ShowURL)//给一组加中间件
	{
		midware.GET("/show", controller.MidwareController{}.Show)

		//配置中间件，最后一个方法是业务处理，之前的都是中间件
		midware.GET("/firsttest", midware2.Midware{}.Printa, controller.MidwareController{}.End)

		midware.GET("/findip", midware2.Midware{}.FindIp, controller.MidwareController{}.End)

		midware.GET("/totaltime", midware2.Midware{}.TotalTime, controller.MidwareController{}.End)

		//中间件 和 控制器之间 互传数据 Get Set
		midware.GET("/transdata", midware2.Midware{}.GetIp, controller.MidwareController{}.ShowData)



	}
}
