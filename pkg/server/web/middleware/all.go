package middleware

import (
	"github.com/gin-gonic/gin"
	"shahaohuo.com/shahaohuo/pkg/server/web/token"
	"strings"
)

const prefix = "Bearer "

func ContentSet() gin.HandlerFunc {
	return func(c *gin.Context) {
		if ok := setUserIdByAuthHeader(c); !ok {
			setUserIdByCookie2Content(c)
		}
	}
}

func setUserIdByCookie2Content(c *gin.Context) bool {
	if _, ok := c.Get(RequestUserId); !ok {
		if value, e := c.Cookie("token"); e == nil {
			if len(value) > 0 {
				if id, ok := token.Verify(value); ok {
					c.Set(RequestUserId, id)
					return true
				}
			}
		}
	}
	return false
}

func setUserIdByAuthHeader(c *gin.Context) bool {
	if _, ok := c.Get(RequestUserId); !ok {
		headerValue := c.GetHeader("Authorization")
		if len(headerValue) > 0 && strings.HasPrefix(headerValue, prefix) {
			if id, ok := token.Verify(headerValue[len(prefix):]); ok {
				c.Set(RequestUserId, id)
				return true
			}
		}
	}
	return false
}
