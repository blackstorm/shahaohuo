package middleware

import (
	"github.com/gin-gonic/gin"
)

const RequestUserId = "request-user-id"

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get(RequestUserId); !ok {
			abort(c)
		}
	}
}
