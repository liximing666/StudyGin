package midware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Midware struct {

}

func (midware Midware) Printa(context *gin.Context)  {
	fmt.Println("aaa")
}

func (midware Midware) FindIp(context *gin.Context)  {
	ip, _ := context.RemoteIP()
	fmt.Println(ip)
}

func (midware Midware) Grope(context *gin.Context)  {

	fmt.Println("group midware")
}

func (midware Midware) TotalTime(context *gin.Context)  {

	start := time.Now().UnixNano()

	fmt.Println("step1")
	//先执行剩下代码再回来
	context.Next()
	time.Sleep(1 * time.Second)

	fmt.Println("step2")

	end := time.Now().UnixNano()

	fmt.Println(end - start)


}

func (midware Midware) ShowURL(context *gin.Context) {
	fmt.Println(context.Request.Method, context.Request.RemoteAddr, context.Request.URL)
}

func (midware Midware) GetIp(context *gin.Context) {
	context.Set("ip", context.Request.RemoteAddr)
}


