package main

import (
	"Gin/src/sessionUse/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()


	//创建session存储引擎，基于redis
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret1"))
	//配置session中间件
	router.Use(sessions.Sessions("mysession1", store))

	//基于redis的session存储
	router.GET("/setredissession", controller.SessionController{}.SetRedisSession)
	router.GET("/getredissession", controller.SessionController{}.GetRedisSession)
	router.GET("/delredissession", controller.SessionController{}.DelRedisSession)




	router.Run(":8888")
}
