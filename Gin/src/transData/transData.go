package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/**/*")//自己导入模板所在的目录及文件



	//获取动态路由的传值 Param
	engine.GET("/dynamic/:id", func(context *gin.Context) {
		s := context.Param("id")
		i, _ := strconv.Atoi(s)

		context.JSON(200, map[string]int{"id": i})
	})


	//get传值 Query
	engine.GET("/get", func(context *gin.Context) {
		s := context.Query("id")
		age := context.DefaultQuery("age", "1")

		context.String(200, "%v,%v", s, age)

	})

	//post传值 PostForm
	engine.GET("/showpost", func(context *gin.Context) {
		context.HTML(200, "posttest.html", gin.H{})
	})


	engine.POST("/post", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.PostForm("age")
		password := context.DefaultPostForm("password", "123456")

		fmt.Println(name, age)
		context.JSON(200, map[string]interface{}{
			"name": name,
			"age": age,
			"password": password,
		})

	})


	//把数据绑定到结构体 ShouldBind

	engine.GET("/tos", func(context *gin.Context) {
		user := User{}
		context.ShouldBind(&user)
		fmt.Println(user)
		context.JSON(200, user)
	})




	engine.GET("/tostruct", func(context *gin.Context) {
		context.HTML(200, "tostruct.html", gin.H{})
	})

	engine.POST("/dostruct", func(context *gin.Context) {

		user := User{}
		context.ShouldBind(&user)
		fmt.Println(user)
		context.JSON(200, user)

		/* 手动
		name := context.PostForm("name")
			age := context.PostForm("age")

			user := User{Name: name, Age: age}
			context.JSON(200, user)
		*/



	})

	engine.Run(":8888")
}
