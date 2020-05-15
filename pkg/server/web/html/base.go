package html

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"shahaohuo.com/shahaohuo/pkg/server/web/middleware"
	"strconv"
)

func notFoundPage(c *gin.Context) {
	c.HTML(404, "404.html", nil)
}

func errorPage(c *gin.Context) {
	c.HTML(500, "500.html", nil)
}

func home302(c *gin.Context) {
	c.Redirect(302, "/")
}

func badRequestPage(c *gin.Context) {
	c.HTML(400, "bad.html", nil)
}

func mustPage(c *gin.Context) int {
	page := c.Query("p")
	i := 1
	if len(page) > 0 {
		x, e := strconv.Atoi(page)
		if e != nil {
			logrus.Warn("bad page query {} ", page)
			return i
		}
		if x <= 0 {
			i = 1
		} else {
			i = x
		}
	}
	return i
}

func mustGetUserIdByContent(c *gin.Context) (string, bool) {
	if id, ok := c.Get(middleware.RequestUserId); ok {
		return id.(string), true
	}

	c.Data(http.StatusForbidden, "text/palin", nil)
	return "", false
}
