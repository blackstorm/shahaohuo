package html

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shahaohuo.com/shahaohuo/pkg/server/web/middleware"
)

func notFoundPage(c *gin.Context) {
	c.HTML(404, "404.html", nil)
}

func errorPage(c *gin.Context) {
	c.HTML(500, "500.html", nil)
}

func mustGetUserIdByContent(c *gin.Context) (string, bool) {
	if id, ok := c.Get(middleware.RequestUserId); ok {
		return id.(string), true
	}

	c.Data(http.StatusForbidden, "text/palin", nil)
	return "", false
}
