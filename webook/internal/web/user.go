package web

import "github.com/gin-gonic/gin"

// UserHandler 我准备在这上面定义所有和用户相关的路由
type UserHandler struct{

}



// 分组注册
func (u *UserHandler)RegisterRoutes(server *gin.Engine){
	ug:=server.Group("/users")
	ug.POST("/user/signup",u.SignUp)
	ug.POST("user/login",u.Login)
	ug.POST("user/edit",u.Edit)
	ug.GET("/user/profile",u.Profile)
}



func (u *UserHandler) SignUp(ctx *gin.Context){

}

func (u *UserHandler) Login(ctx *gin.Context){

}

func (u *UserHandler) Edit (ctx *gin.Context){

}

func (u *UserHandler) Profile(ctx *gin.Context){

}