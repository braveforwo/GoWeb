package routers

import (
	"GoWeb/config/middleware"
	"GoWeb/domain"
	"GoWeb/server/impl"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadServer(e *gin.Engine) {
	e.POST("/registerService", registerServiceHandler)
	e.POST("/loginService", loginServiceHandler)
}

func registerServiceHandler(c *gin.Context) {
	//session := sessions.Default(c)
	var json domain.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h := md5.New()
	h.Write([]byte(json.Password))
	cipherStr := h.Sum(nil)
	json.Password = hex.EncodeToString(cipherStr)
	//if session.Get("hello") != "world" {
	//	session.Set("hello", "world")
	//	//记着调用save方法，写入session
	//	session.Save()
	//}
	resgister := impl.ResgisterServiceImpl{}
	if err := resgister.Register(&json); err != nil {
		middleware.LogerMiddlewareCom(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	middleware.LogerMiddlewareCom("注册成功")
	fmt.Println(json)
	//c.JSON(200, gin.H{"hello": session.Get("hello")})
	c.JSON(http.StatusBadRequest, gin.H{"msg": "注册成功"})
}

func loginServiceHandler(c *gin.Context) {
	var json domain.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	login := impl.LoginServiceImpl{}
	if err := login.Login(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	session := sessions.Default(c)
	session.Set("user", json)
	session.Save()
	c.JSON(200, gin.H{"msg": "登录成功！"})
}
