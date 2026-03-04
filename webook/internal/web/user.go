package web

import (
	"net/http"

	"github.com/Powder217/go_practice/webook/internal/domain"
	"github.com/Powder217/go_practice/webook/internal/service"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

// 正则表达式
const (
	emailregex    = `^[\w.-]+@[\w.-]+\.[a-zA-Z]{2,6}$`
	passwordregex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*])?[A-Za-z\d!@#$%^&*]{8,}$`
)	

// UserHandler 我准备在这上面定义所有和用户相关的路由
type UserHandler struct{
	emailRegex *regexp2.Regexp
	passwordRegex *regexp2.Regexp
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService)*UserHandler{
	return &UserHandler{
		svc:svc,
		emailRegex: regexp2.MustCompile(emailregex,regexp2.None),
		passwordRegex: regexp2.MustCompile(passwordregex,regexp2.None),
	}
}

// 分组注册
func (u *UserHandler)RegisterRoutes(server *gin.Engine){
	ug:=server.Group("/users")
	ug.POST("/signup",u.SignUp)
	ug.POST("/login",u.Login)
	ug.POST("/edit",u.Edit)
	ug.GET("/profile",u.Profile)
}

func (u *UserHandler) SignUp(ctx *gin.Context){

	// 内部结构体
	type SignUpReq struct{
		Email string `json:"email"` 
		ConfirmPassWord string `json:"confirmPassword"`
		PassWord string `json:"password"`
	}
	var req SignUpReq
	// Bind会根据方法和Content-Type自动选择绑定引擎，根据“Content-Type”头部信息，会使用不同的绑定方式，例如：
	// "application/json" --> JSON绑定
	// "application/xml" --> XML绑定
	// 它根据Content-Type（例如JSON或XML）解析请求的正文。它将有效负载解码为指定为指针的结构体。如果输入无效，则写入400错误并在响应中设置Content-Type头为“text/plain”。

	// 确保你正确的拿到了数据
	if err:=ctx.Bind(&req);err !=nil{
		return
	}

	// 注册校验

	u.emailRegex=regexp2.MustCompile(emailregex,regexp2.None)
	u.passwordRegex=regexp2.MustCompile(passwordregex,regexp2.None)
	if req.ConfirmPassWord!=req.PassWord{
		ctx.String(http.StatusOK,"两次密码输入不一致")
		return
	}
	ok,err:=u.emailRegex.MatchString(req.Email)
	if err!=nil{
		ctx.String(http.StatusOK,"系统错误")
		return 
	}
	if !ok{
		ctx.String(http.StatusOK,"邮箱格式错误")
		return
	}
	ok,err=u.passwordRegex.MatchString(req.PassWord)
	if err!=nil{
		ctx.String(http.StatusOK,"系统错误")
		return
	}
	if !ok{
		ctx.String(http.StatusOK,"密码格式有误,密码长度必须大于8位且包含数字、特殊字符、大小写字母")
		return
	}
	err=u.svc.SignUp(ctx,domain.User{
		Email: req.Email,
		Password: req.PassWord,
	})
	if err !=nil{
		ctx.String(http.StatusOK,"系统错误")
		return 
	}

	ctx.String(http.StatusOK,"注册成功")

}	

func (u *UserHandler) Login(ctx *gin.Context){

}

func (u *UserHandler) Edit (ctx *gin.Context){

}

func (u *UserHandler) Profile(ctx *gin.Context){

}