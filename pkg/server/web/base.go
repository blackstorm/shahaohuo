package web

import (
	"github.com/gin-gonic/gin"
	"shahaohuo.com/shahaohuo/pkg/server/web/middleware"
)

type UriIdParams struct {
	Id string `uri:"id" binding:"required,max=32,min=1"`
}

func GetUserIdByContent(c *gin.Context) (string, bool) {
	if id, ok := c.Get(middleware.RequestUserId); ok {
		return id.(string), true
	}

	return "", false
}
