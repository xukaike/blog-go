package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"blog/pkg/ctx"
)

func Auth(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("userInfo")
	if user == nil{
		ctx.Response(c,http.StatusUnauthorized,ctx.UserNotLogin,nil)
		return
	}
	c.Set("userInfo",user)
	c.Next()
}