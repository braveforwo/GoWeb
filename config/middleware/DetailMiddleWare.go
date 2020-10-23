package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func DetailMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/detail/assert") {
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, "/detail/", "", 1)
			c.Redirect(http.StatusTemporaryRedirect, strings.Replace(c.Request.URL.Path, "/detail/", "", 1))
		}
	}
}
