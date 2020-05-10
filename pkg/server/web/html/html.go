package html

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/service"
	"shahaohuo.com/shahaohuo/pkg/server/web"
	"time"
)

var shanghai, _ = time.LoadLocation("Asia/Shanghai")

func Index(ctx *gin.Context) {
	hs, e := service.FindUserBusinessHaohuosByLimit(4)
	if e == nil {
		res := gin.H{"NewHaohuos": hs}
		mostFavorite, e := queryMostFavorite()
		if e == nil {
			res["MostFavorite"] = mostFavorite
			us, e := service.FindLatestUsersByLimit(64)
			if e == nil {
				res["LatestUsers"] = us
				res["Statics"] = service.Statics()
				ctx.HTML(http.StatusOK, "index.html", res)
				return
			}
		}
	}
	errorPage(ctx)
}

func queryMostFavorite() ([]orm.BusinessHaohuo, error) {
	now := time.Now().In(shanghai)
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return service.FindMostFavoriteBusinessHaohuosByDateAndLimit(start, now, 4)
}

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func Logout(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "logout.html", nil)
}

func Share(c *gin.Context) {
	tags, _ := orm.FindAllTag()
	c.HTML(http.StatusOK, "share.html", gin.H{"Now": time.Now().In(shanghai), "Tags": tags})
}

func Users(c *gin.Context) {
	var uriParams web.UriIdParams
	if e := c.BindUri(&uriParams); e != nil {
		// TODO bad page
		errorPage(c)
		return
	}
	u, e := service.FindBasicUserById(uriParams.Id)
	if e == nil {
		if u == nil {
			notFoundPage(c)
			return
		}
		ret := gin.H{
			"User": u,
		}
		hs, e := service.FindUserBusinessHaohuosByUserIdAndLimit(uriParams.Id, 24)
		comments, _ := orm.FindUserCommentsByUserId(uriParams.Id, 10)
		if e == nil {
			ret["Haohuos"] = hs
			ret["Comments"] = comments
			c.HTML(http.StatusOK, "user.html", ret)
			return
		}
	}
	errorPage(c)
}

func Settings(c *gin.Context) {
	if uid, ok := mustGetUserIdByContent(c); ok {
		if u, e := service.FindUserById(uid); e == nil {
			c.HTML(http.StatusOK, "settings.html", gin.H{"User": u})
			return
		}
		errorPage(c)
	}

}
