package main

import (
	"strings"
	"time"

	"github.com/Powder217/go_practice/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func main()  {
	server:=gin.Default()
	server.Use(cors.New(cors.Config{
	// 允许跨域的域名 
	AllowOrigins:     []string{"http://localhost:3000"},
	// 允许的请求方法 不填则默认是全部方法
    // AllowMethods:     []string{"PUT", "PATCH"},
	// 允许的那些请求头可以被前端代码访问
    AllowHeaders:     []string{"authorization","content-type"},
	// 允许哪些响应头可以被前端代码访问
    ExposeHeaders:    []string{"Content-Length"},
	// 是否允许携带 cookies
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      if strings.HasPrefix(origin,"http://localhost"){
	  // 如果访问的origin 的前缀包含 http://localhost 则允许访问
	  return true
	  }
	  return strings.Contains(origin,"")
    },
    MaxAge: 12 * time.Hour,
	}))
	u:=&web.UserHandler{}
	u.RegisterRoutes(server)
	server.Run(":8080")
}