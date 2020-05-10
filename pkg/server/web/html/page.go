package html

import "github.com/gin-gonic/gin"

func PageHandle(page string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, page+".html", nil)
	}
}
