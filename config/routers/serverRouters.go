package routers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoadServer(e *gin.Engine) {
	e.GET("/test", testHandler)
	e.GET("/gettest", gettestHandler)
}

func testHandler(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("hello") != "world" {
		session.Set("hello", "world")
		//记着调用save方法，写入session
		session.Save()
	}
	c.JSON(200, gin.H{"hello": session.Get("hello")})
}

func gettestHandler(c *gin.Context) {
	session := sessions.Default(c)
	c.Next()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	c.JSON(200, gin.H{"hello": session.Get("hello")})
}
