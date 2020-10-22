package routers

import (
	"GoWeb/config/middleware"
	"GoWeb/domain"
	"GoWeb/server/impl"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadHtml(e *gin.Engine) {
	e.GET("/index", middleware.AuthenticationMiddleware(), indexHandler)
	e.GET("/article", middleware.AuthenticationMiddleware(), articleHandler)
	e.GET("/detail", middleware.AuthenticationMiddleware(), detailHandler)
	e.GET("/about", middleware.AuthenticationMiddleware(), aboutHandler)
	e.GET("/resource", middleware.AuthenticationMiddleware(), resourceHandler)
	e.GET("/timeline", middleware.AuthenticationMiddleware(), timelineHandler)
	e.GET("/login", loginHandler)
	e.GET("/register", registerHandler)
	e.GET("/404", ErrorHandler)
	e.GET("/errors", ErrorsHandler)
	e.GET("/markdown", markdownHandler)
	e.POST("/articleList", articleListHandler)
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

func ErrorHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{
		"title": "Main website",
	})
}

func ErrorsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "errors.html", gin.H{
		"title": "Main website",
	})
}
func markdownHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "markdown.html", gin.H{
		"title": "Main website",
	})
}

func articleListHandler(c *gin.Context) {
	var articleSearchCondition domain.ArticleSearchCondition
	if err := c.ShouldBind(&articleSearchCondition); err != nil {
		fmt.Println(err)
	}
	searchArticleServiceImpl := impl.SearchArticleServiceImpl{}
	err, articlelist := searchArticleServiceImpl.SearchArticleService(&articleSearchCondition)
	if err != nil {
		c.HTML(http.StatusOK, "articlelist.html", articlelist)
		return
	}
	c.HTML(http.StatusOK, "articlelist.html", articlelist)
}
