package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name string
	Age int
}


// 响应 各种类型的数据
func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/**/*")//自己导入模板所在的目录

	engine.GET("/string", func(context *gin.Context) {
		context.String(http.StatusOK, "返回字符串")
	})
	
	engine.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"a": 1,
			"b": 2,
			"c": 3,
			"d": map[int]int{
				1: 6,
			},
		})
	})

	engine.GET("/struct", func(context *gin.Context) {
		context.JSON(http.StatusOK, Person{Name: "aaa", Age: 16})
	})

	engine.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, Person{Name: "ccc", Age: 16})
	})

	//跳转到一个页面，并且把数据传到前端
	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", map[string]interface{}{"age": 16})
	})





	engine.Run()

}
