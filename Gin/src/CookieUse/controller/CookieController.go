package controller

import "github.com/gin-gonic/gin"

type CookieController struct {

}

func (c CookieController) ToLogin(context *gin.Context) {

	//先验证cookie
	username, _ := context.Cookie("username")

	if username == "" {
		context.HTML(200, "login.html", gin.H{})
	}else {
		context.HTML(200, "index.html", gin.H{"username": username})
	}

}

func (c CookieController) ToRegister(context *gin.Context)  {
	context.HTML(200, "register.html", gin.H{})
}

func (c CookieController) DoRegister(context *gin.Context) {
	username := context.PostForm("username")


	context.JSON(200, gin.H{"username": username})

}

func (c CookieController) DoLogin(context *gin.Context) {

	username := context.PostForm("username")
	rem := context.PostForm("rem")

	//保存cookie path 保存路径 domain作用域
	//cookie 存结构体对象 先转化为 json 或者 map 存
	//设置domain在 不同的 二级域名中共享cookie

	if rem == "on" {
		context.SetCookie("username", username, 3600, "", ".lxm.com",false, true)
	}

	context.String(200,"login ok")
}

func (c CookieController) LoginOut(context *gin.Context) {

	//ttl改成-1表示删除cookie
	context.SetCookie("username", "", -1, "", "",true, true)
	context.HTML(200, "login.html", gin.H{})


}


