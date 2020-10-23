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
	e.GET("/index", indexHandler)
	e.GET("/article", articleHandler)
	e.GET("/detail", detailHandler)
	e.GET("/about", middleware.AuthenticationMiddleware(), aboutHandler)
	e.GET("/resource", middleware.AuthenticationMiddleware(), resourceHandler)
	e.GET("/timeline", middleware.AuthenticationMiddleware(), timelineHandler)
	e.GET("/login", loginHandler)
	e.GET("/register", registerHandler)
	e.GET("/404", ErrorHandler)
	e.GET("/errors", ErrorsHandler)
	e.GET("/markdown", middleware.AuthenticationMiddleware(), markdownHandler)
	e.POST("/articleList", articleListHandler)
	e.GET("/modifyInformation", modifyInformationHandler)
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
	var elasticArticleModel domain.ElasticArticleModel
	if err := c.ShouldBind(&elasticArticleModel); err != nil {
		c.HTML(http.StatusOK, "detail.html", gin.H{"article": nil, "msg": err})
		return
	}
	var article domain.Article
	var err error
	searchArticleService := impl.SearchArticleServiceImpl{}
	if err, article = searchArticleService.UpdateArticlePageViewsToMysql(&elasticArticleModel); err != nil {
		fmt.Println(err)
		return
	}
	elasticArticleModel.Pageviews = article.Pageviews
	if err = searchArticleService.UpdateArticlePageViewsToElastic(&elasticArticleModel); err != nil {
		fmt.Println(err)
		return
	}
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"article": article,
		"msg":     err,
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
	fmt.Println(articleSearchCondition)
	searchArticleServiceImpl := impl.SearchArticleServiceImpl{}
	err, articlelist := searchArticleServiceImpl.SearchArticleServiceFromElastic(&articleSearchCondition)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "articlelist.html", gin.H{"articlelist": articlelist, "articleSearchCondition": articleSearchCondition, "alert": err})
		return
	}
	c.HTML(http.StatusOK, "articlelist.html", gin.H{"articlelist": articlelist, "articleSearchCondition": articleSearchCondition})
}
func modifyInformationHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "modifyInformation.html", gin.H{
		"title": "Main website",
	})
}
