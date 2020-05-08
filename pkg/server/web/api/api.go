package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"shahaohuo.com/shahaohuo/pkg/server/dto"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/service"
	"shahaohuo.com/shahaohuo/pkg/server/storage"
	"shahaohuo.com/shahaohuo/pkg/server/web"
	"shahaohuo.com/shahaohuo/pkg/tool"
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
		if h, e := service.CreateOrUpdate(&req, userId); e == nil {
			ret := &CreateOrUpdateResponse{}
			_ = copier.Copy(ret, h)
			ok(c, ret)
		} else {
			if e == service.ResourcesExist {
				bad(c, "haohuo exist")
				return
			}
			error(c, "server error")
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
		error(c, newBadResp(-1, "server error"))
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
		error(c, newBadResp(-1, "server error"))
	}
}

func UploadImage(c *gin.Context) {
	userId, _ := mustGetUserIdByContent(c)
	fileHeader, err := c.FormFile("image")

	if err != nil {
		logrus.Error(err)
		bad(c, "get file error")
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		logrus.Error(err)
		error(c, "server error")
		return
	}

	// TODO zero copy
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Error(err)
		error(c, "server error")
		return
	}

	contentType := http.DetectContentType(bytes)
	if !tool.CheckIsContain(contentType, tool.PNG, tool.JPEG) {
		bad(c, newBadResp(1, "file must is an image"))
		return
	}

	id := xid.New().String()
	imagePath := "/haohuos/images/" + id
	if err := storage.GetBucket().Upload(imagePath, contentType, bytes); err != nil {
		logrus.Error(err)
		error(c, "server error")
		return
	}

	go orm.SaveImage(userId, id, imagePath)

	ok(c, gin.H{
		"path": imagePath,
	})
}
