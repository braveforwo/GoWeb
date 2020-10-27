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
	"time"
)

func LoadServer(e *gin.Engine) {
	e.POST("/registerService", registerServiceHandler)
	e.POST("/loginService", loginServiceHandler)
	e.POST("/upload", middleware.AuthenticationMiddleware(), uploadServiceHandler)
	e.POST("/uploadArticle", middleware.AuthenticationMiddleware(), uploadArticleServiceHandler)
	e.POST("/getUserMessage", getUserMessageServiceHandler)
	e.POST("/logOut", middleware.AuthenticationMiddleware(), logOutServiceHandler)
	e.POST("/commitComment", middleware.AuthenticationMiddleware(), commitCommentServiceHandler)
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
	//fmt.Println(json)
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

func uploadServiceHandler(c *gin.Context) {
	file, _ := c.FormFile("editormd-image-file")
	// 上传文件到指定的路径
	err := c.SaveUploadedFile(file, "assert/uploadImages/"+file.Filename)
	if err == nil {
		responseMessage := domain.UploadResponseModel{}
		responseMessage.Success = 1
		responseMessage.Message = "上传成功"
		responseMessage.Url = "assert/uploadImages/" + file.Filename
		c.JSON(200, responseMessage)
		return
	}
	responseMessage := domain.UploadResponseModel{}
	responseMessage.Success = 0
	responseMessage.Message = "上传失败"
	fmt.Println(err)
	c.JSON(200, responseMessage)
}

func uploadArticleServiceHandler(c *gin.Context) {
	var article domain.Article
	if err := c.Bind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	var user domain.User = session.Get("user").(domain.User)
	article.UserId = user.Id
	article.Pageviews = 0
	t := time.Now()
	timestr := t.Format("2006-01-02 15:04:05")
	article.Time = timestr
	uploadArticleServiceImpl := impl.UploadArticleServiceImpl{}
	if err := uploadArticleServiceImpl.UploadArticle(&article); err != nil {
		c.JSON(200, gin.H{"msg": err})
		return
	}
	if err := uploadArticleServiceImpl.UploadArticleToElastic(&article); err != nil {
		c.JSON(200, gin.H{"msg": err})
		return
	}

	//fmt.Println(article)
	c.JSON(200, gin.H{"msg": "上传成功！"})
}

func getUserMessageServiceHandler(c *gin.Context) {
	session := sessions.Default(c)
	var user domain.User
	if session.Get("user") != nil {
		user = session.Get("user").(domain.User)
	}
	if user.UserName != "" {
		c.JSON(200, gin.H{"msg": user.UserName})
		return
	}
	c.JSON(400, gin.H{"msg": "获取个人信息失败！"})
}

func logOutServiceHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{"msg": "成功退出登录！"})
}

func commitCommentServiceHandler(c *gin.Context) {
	var comment domain.Comment
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(200, gin.H{"msg": "参数错误！"})
		return
	}
	session := sessions.Default(c)
	user := session.Get("user").(domain.User)
	comment.Userid = user.Id
	t := time.Now()
	timestr := t.Format("2006-01-02 15:04:05")
	comment.Time = timestr
	commentService := impl.CommentServiceImpl{}
	if err := commentService.CreateComment(&comment); err != nil {
		c.JSON(200, gin.H{"msg": "提交失败！"})
		return
	}
	var article domain.Article
	article.Id = comment.Articleid
	if err := commentService.UpdateCommentNumToMysql(&article); err != nil {
		c.JSON(200, gin.H{"msg": "提交失败！"})
		return
	}
	if err := commentService.UpdateCommentNumToElastic(&article); err != nil {
		c.JSON(200, gin.H{"msg": "提交失败！"})
		return
	}
	c.JSON(200, gin.H{"msg": "提交成功！"})
}
