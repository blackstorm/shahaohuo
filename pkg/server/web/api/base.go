package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shahaohuo.com/shahaohuo/pkg/server/web/middleware"
)

type BadResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newBadResp(code int, message string) *BadResponse {
	return &BadResponse{
		Code:    code,
		Message: message,
	}
}

func bad(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusBadRequest, obj)
}

func notfound(c *gin.Context) {
	c.JSON(http.StatusBadRequest, nil)
}

func ok(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func internalServerError(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusInternalServerError, obj)
}

func mustGetUserIdByContent(c *gin.Context) (string, bool) {
	if id, ok := c.Get(middleware.RequestUserId); ok {
		return id.(string), true
	}

	c.JSON(401, &BadResponse{
		Code:    1,
		Message: "failed token",
	})
	return "", false
}
