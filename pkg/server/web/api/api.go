package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"shahaohuo.com/shahaohuo/pkg/server/dto"
	"shahaohuo.com/shahaohuo/pkg/server/image"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/service"
	"shahaohuo.com/shahaohuo/pkg/server/web"
	"time"
)

type CreateOrUpdateResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"imageUrl"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func CreateOrUpdateHaohuo(c *gin.Context) {
	var req dto.HaohuoRequest
	var uriParams web.UriIdParams

	if e := c.BindUri(&uriParams); e != nil {
		bad(c, gin.H{"errors": e.Error()})
		return
	}

	if e := c.ShouldBindJSON(&req); e != nil {
		bad(c, gin.H{"errors": e.Error()})
		return
	}

	if userId, get := mustGetUserIdByContent(c); get {
		req.Id = uriParams.Id

		if !image.CheckIsUserImage(userId, req.ImageUrl) {
			bad(c, "bad image url")
			return
		}

		if h, e := service.CreateOrUpdate(&req, userId); e == nil {
			ret := &CreateOrUpdateResponse{}
			_ = copier.Copy(ret, h)
			ok(c, ret)
		} else {
			if e == service.ResourcesExist {
				bad(c, "haohuo exist")
				return
			}
			internalServerError(c, "server error")
		}
	}
}

func FavoriteHaohuo(c *gin.Context) {
	var uriParams web.UriIdParams
	if e := c.BindUri(&uriParams); e != nil {
		bad(c, gin.H{"errors": e.Error()})
		return
	}

	if userId, get := mustGetUserIdByContent(c); get {
		e := service.HaohuoFavorite(uriParams.Id, userId)
		if e == nil {
			ret, e := orm.CountFavoritesByHaohuoId(uriParams.Id)
			if e == nil {
				ok(c, ret)
				return
			}
		}
		internalServerError(c, newBadResp(-1, "server error"))
	}
}

func CommentHaohuo(c *gin.Context) {
	var req dto.CommentRequest
	var uriParams web.UriIdParams

	if e := c.BindUri(&uriParams); e != nil {
		bad(c, gin.H{"errors": e.Error()})
		return
	}

	if e := c.ShouldBindJSON(&req); e != nil {
		bad(c, gin.H{"errors": e.Error()})
		return
	}

	if userId, get := mustGetUserIdByContent(c); get {
		comment, e := service.CommentHaohuo(uriParams.Id, userId, req)
		if e == service.ResourcesNotFound {
			notfound(c)
			return
		}
		if e == nil {
			var ret dto.CommentResponse
			_ = copier.Copy(&ret, comment)
			ok(c, &ret)
			return
		}
		internalServerError(c, newBadResp(-1, "server error"))
	}
}
