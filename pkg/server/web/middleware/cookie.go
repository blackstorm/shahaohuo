package middleware

import (
	"github.com/gin-gonic/gin"
)

func CookieAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get(RequestUserId); !ok {
			abortWithRedirect2Login(c)
		}
	}
}
