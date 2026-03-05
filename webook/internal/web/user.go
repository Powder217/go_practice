package web

import (
	"net/http"

	"github.com/Powder217/go_practice/webook/internal/domain"
	"github.com/Powder217/go_practice/webook/internal/service"
	"github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 正则表达式
const (
	emailregex    = `^[\w.-]+@[\w.-]+\.[a-zA-Z]{2,6}$`
	passwordregex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*])?[A-Za-z\d!@#$%^&*]{8,}$`
)	

type UserHandler struct{
	emailRegex *regexp2.Regexp
	passwordRegex *regexp2.Regexp
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService)*UserHandler{
	return &UserHandler{
		emailRegex: regexp2.MustCompile(emailregex,regexp2.None),
		passwordRegex: regexp2.MustCompile(passwordregex,regexp2.None),
		svc:svc,
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

func (h *UserHandler) SignUp(ctx *gin.Context){

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

	h.emailRegex=regexp2.MustCompile(emailregex,regexp2.None)
	h.passwordRegex=regexp2.MustCompile(passwordregex,regexp2.None)
	if req.ConfirmPassWord!=req.PassWord{
		ctx.String(http.StatusOK,"两次密码输入不一致")
		return
	}
	ok,err:=h.emailRegex.MatchString(req.Email)
	if err!=nil{
		ctx.String(http.StatusOK,"系统错误")
		return 
	}
	if !ok{
		ctx.String(http.StatusOK,"邮箱格式错误")
		return
	}
	ok,err=h.passwordRegex.MatchString(req.PassWord)
	if err!=nil{
		ctx.String(http.StatusOK,"系统错误")
		return
	}
	if !ok{
		ctx.String(http.StatusOK,"密码格式有误,密码长度必须大于8位且包含数字、特殊字符、大小写字母")
		return
	}
	err=h.svc.SignUp(ctx,domain.User{
		Email: req.Email,
		PassWord: req.PassWord,
	})
	if err ==service.ErrUserDuplicateEmail{
		ctx.String(http.StatusOK,"邮箱已存在")
		return 
	}

	ctx.String(http.StatusOK,"注册成功")

}	

func (h *UserHandler) Login(ctx *gin.Context){
	type LoginReq struct{
		Email string `json:"email"`
		PassWord string `json:"password"`
	}

	var req LoginReq
	if err:=ctx.Bind(&req);err!=nil{
		return
	}
	u,err:=h.svc.Login(ctx,req.Email,req.PassWord)

	if err==service.ErrInvalidUserOrPassword{
		ctx.String(http.StatusOK,"用户名或密码错误")
		return 
	}
	if err!=nil{
		ctx.String(http.StatusOK,"系统错误")
		return
	}

	// 设置session
	session:=sessions.Default(ctx)
	// 可以处理session的各个值	
	session.Set("userId",u.Id)
	session.Save()
	ctx.String(http.StatusOK,"登陆成功")


}

func (u *UserHandler) Edit (ctx *gin.Context){

}

func (u *UserHandler) Profile(ctx *gin.Context){
	ctx.String(http.StatusOK,"profile")
}