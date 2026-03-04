package main

import (
	"strings"
	"time"

	"github.com/Powder217/go_practice/webook/internal/repository"
	"github.com/Powder217/go_practice/webook/internal/repository/dao"
	"github.com/Powder217/go_practice/webook/internal/service"
	"github.com/Powder217/go_practice/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 关于数据链路的理解
// Web → Service → Repository  (业务逻辑层)
// DAO                        (数据库操作层)
// Domain                     (连接两者的领域模型)
// Domain 到 DAO 的转换由 Repository 来完成，Repository 负责将领域模型转换为 DAO 模型，并调用 DAO 层的方法来执行数据库操作。（creat 方法）

func main()  {

	db:=initDB()
	u:=initUser(db)
	server:=initWebServer()
	u.RegisterRoutes(server)
	server.Run(":8080")
}

func initDB()*gorm.DB{
	db,err:=gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err!=nil{
		panic(err)
		// 我只会在初始化中 panic，发生panic 代表着 goroutine 的结束
	}
	return db
}

func initUser(db *gorm.DB)*web.UserHandler{
	ud:=dao.NewUserDAO(db)
	repo:=repository.NewUserRepository(ud)
	svc:=service.NewUserService(&repo)
	u:=web.NewUserHandler(svc)

	err:=dao.InitTable(db)
	if err != nil {
			panic(err)
	}
	return u
}

func initWebServer()*gin.Engine{
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
	return server
}