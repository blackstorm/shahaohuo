package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/server/dto"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/service"
	"shahaohuo.com/shahaohuo/pkg/server/web"
)

func AddHaoHuoVideo(c *gin.Context) {
	var uriParams web.UriIdParams
	if e := c.BindUri(&uriParams); e != nil {
		bad(c, gin.H{"errors": e.Error()})
		return
	}

	var req dto.HaohuoVideoRequest
	if e := c.ShouldBindJSON(&req); e != nil {
		bad(c, gin.H{"errors": e.Error()})
		return
	}

	haohuo, e := service.FindHaohuoById(uriParams.Id)
	if e != nil {
		if service.IsResourcesNotFound(e) {
			notfound(c)
			return
		}
	}

	if u, ok := mustGetUserIdByContent(c); ok {
		v := orm.NewBilibiliIframeVideo(u, haohuo.Id, req.Url)
		if e := v.Create(); e != nil {
			logrus.Error(e)
			internalServerError(c, nil)
			return
		}
		c.JSON(201, nil)
	}

}
