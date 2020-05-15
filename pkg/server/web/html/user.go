package html

import (
	"github.com/gin-gonic/gin"
	p "shahaohuo.com/shahaohuo/pkg/page"
	"shahaohuo.com/shahaohuo/pkg/request"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/service"
)

func UserHaohuosPage(c *gin.Context) {
	var params request.StrIDUriParams
	if e := c.BindUri(&params); e != nil {
		badRequestPage(c)
		return
	}
	_userHaohuoPages(c, params.Id, mustPage(c))
}

func _userHaohuoPages(c *gin.Context, userId string, page int) {
	if page <= 0 {
		page = 1
	}

	u, e := service.FindUserById(userId)
	if e != nil {
		errorPage(c)
		return
	}
	if u == nil {
		notFoundPage(c)
		return
	}

	counts := orm.CountHaohuosByUserId(userId)

	hs, e := service.FindUserBusinessHaohuosByUserIdAndPage(userId, page, 8)
	if e != nil {
		errorPage(c)
		return
	}

	baseUrl := "/users/" + userId + "/haohuos"
	c.HTML(200, "user_haohuos.html", gin.H{
		"Page":    p.NewPage(8, page, counts, 6, baseUrl),
		"Haohuos": hs,
		"User":    u,
	})

}
