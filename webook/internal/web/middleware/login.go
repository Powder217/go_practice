package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddleWareBuilder struct{

}

func NewLoginMiddleWareBuilder()*LoginMiddleWareBuilder{
	return &LoginMiddleWareBuilder{}

}

func (l *LoginMiddleWareBuilder) CheckLogin() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// 不需要校验 
		if ctx.Request.URL.Path=="/users/login" || ctx.Request.URL.Path=="/users/signup"{
			return 
		}

		session:=sessions.Default(ctx)
		if session==nil{
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return 
		}
		id:=session.Get("userId")
		if id==nil{
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return 
		}
	}
}