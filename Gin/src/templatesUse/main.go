package main

import (
	"Gin/src/templatesUse/tools"
	"github.com/gin-gonic/gin"
	template2 "html/template"
	"net/http"
	"text/template"
)

type Person struct {
	Name string
	Age int
}

//模板的渲染和模板语法
func main() {

	engein := gin.Default()

	//注册模板函数
	engein.SetFuncMap(template2.FuncMap(template.FuncMap{
		"toDate": tools.ToDate,
		"show": tools.Show,
	}))

	engein.LoadHTMLGlob("templates/**/*")//自己导入模板所在的目录及文件

	engein.GET("/admin", func(context *gin.Context) {
		context.HTML(http.StatusOK, "admin.html", map[string]interface{}{
			"name": "admin",
			"userList": []Person{
				{Name: "a", Age: 16},
				{Name: "b", Age: 26},
				{Name: "c", Age: 36},
			},
			"timeStrap": 1063958699,
			"s1": "abc",
			"s2": "abc",
		})
	})

	engein.GET("/else", func(context *gin.Context) {
		context.HTML(http.StatusOK, "admin.html", map[string]interface{}{
			"list": []string{},
		})
	})




	engein.Run(":6666")


}
