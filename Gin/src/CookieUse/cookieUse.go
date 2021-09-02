package main

import (
	"Gin/src/CookieUse/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")


	router.GET("/register", controller.CookieController{}.ToRegister)
	router.POST("/doregister", controller.CookieController{}.DoRegister)
	router.GET("/login", controller.CookieController{}.ToLogin)
	router.POST("/dologin",controller.CookieController{}.DoLogin)
	router.GET("/logout",controller.CookieController{}.LoginOut)

	router.Run(":8888")

}
