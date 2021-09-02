package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionController struct {

}

func (c SessionController) GetCookieSession(context *gin.Context) {
	session := sessions.Default(context)
	s1 := session.Get("session1").(string)
	context.JSON(200, gin.H{"session1": s1})
	
}

func (c SessionController) SetCookieSession(context *gin.Context) {

	//创建session对象
	session := sessions.Default(context)
	session.Set("session1", "session by cookie")
	session.Save()


	
}

func (c SessionController) SetRedisSession(context *gin.Context) {

	session := sessions.Default(context)
	//session的过期时间和相关设置
	session.Options(sessions.Options{MaxAge: 30})
	session.Set("session1",  "by redis")
	session.Save()
	
}

func (c SessionController) GetRedisSession(context *gin.Context) {
	session := sessions.Default(context)
	value := session.Get("session1")
	fmt.Println(value)
}



func (c SessionController) DelRedisSession(context *gin.Context) {
	session := sessions.Default(context)
	session.Options(sessions.Options{MaxAge: -1})
}






