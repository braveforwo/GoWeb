package main

import (
	"GoWeb/config/middleware"
	"GoWeb/config/routers"
	"GoWeb/domain"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	gob.Register(domain.User{})
	store := cookie.NewStore([]byte("secret"))
	r := gin.Default()
	//使用中间件
	r.Use(sessions.Sessions("mysession", store), middleware.LogerMiddleware())
	r.NoRoute(routers.ErrorHandler)
	//设置html路径
	r.LoadHTMLGlob("assert/html/*")
	//设置静态资源路径
	r.StaticFS("assert", http.Dir("./assert"))
	//Server路由
	routers.LoadServer(r)
	//html路由
	routers.LoadHtml(r)
	r.Run(":8080")

}

func formatAsDate(s []string, index int) string {
	return s[index]
}
