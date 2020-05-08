package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func abort(c *gin.Context) {
	c.JSON(401, gin.H{
		"code":    1,
		"message": "failed token",
	})
	c.Abort()
}

func abortAndForbidden(c *gin.Context) {
	c.Data(http.StatusForbidden, "text/plain", nil)
	c.Abort()
}

func abortWithRedirect2Login(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login")
	c.Abort()
}
