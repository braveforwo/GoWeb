package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadHtml(e *gin.Engine) {
	e.GET("/index", indexHandler)
	e.GET("/article", articleHandler)
	e.GET("/detail", detailHandler)
	e.GET("/about", aboutHandler)
	e.GET("/resource", resourceHandler)
	e.GET("/timeline", timelineHandler)
	e.GET("/login", loginHandler)
	e.GET("/register", registerHandler)
	e.GET("/404", errorHandler)
}
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Main website",
	})
}

func articleHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "article.html", gin.H{
		"title": "Main website",
	})
}
func detailHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"title": "Main website",
	})
}
func aboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"title": "Main website",
	})
}
func resourceHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "resource.html", gin.H{
		"title": "Main website",
	})
}
func timelineHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "timeline.html", gin.H{
		"title": "Main website",
	})
}
func loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Main website",
	})
}
func registerHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Main website",
	})
}

func errorHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{
		"title": "Main website",
	})
}
