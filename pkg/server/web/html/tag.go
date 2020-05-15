package html

import (
	"github.com/gin-gonic/gin"
	p "shahaohuo.com/shahaohuo/pkg/page"
	"shahaohuo.com/shahaohuo/pkg/request"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/service"
	"shahaohuo.com/shahaohuo/pkg/util"
	"strconv"
)

const defaultSize = 12
const defaultRenderPageItems = 6

func TagsPage(c *gin.Context) {
	var uriParams request.Int64IDUriParams
	if e := c.BindUri(&uriParams); e != nil {
		badRequestPage(c)
		return
	}
	_tagPages(c, uriParams.Id, mustPage(c))
}

func _tagPages(c *gin.Context, id int64, page int) {
	tag := orm.FindTagById(id)
	if tag == nil {
		notFoundPage(c)
		return
	}

	countsChan := make(chan util.AsyncResult)
	defer close(countsChan)
	go func() {
		countsChan <- util.AsyncResult{
			Ret:   orm.CountHaohuosByTagId(id),
			Error: nil,
		}
	}()

	haohuosChan := make(chan util.AsyncResult)
	defer close(haohuosChan)
	go func() {
		haohuoIds := orm.FindHaohuoIdsByTagId(id, page, defaultSize)
		ret := util.AsyncResult{Ret: nil, Error: nil}
		if len(haohuoIds) > 0 {
			haohuos, e := service.FindUserBusinessHaohuosByIds(haohuoIds)
			ret.Ret = haohuos
			ret.Error = e

		}
		haohuosChan <- ret
	}()

	counts := <-countsChan
	haohuos := <-haohuosChan

	if util.CheckAsyncResultsError(counts, haohuos) {
		errorPage(c)
		return
	}

	c.HTML(200, "tag.html", gin.H{
		"Page":    p.NewPage(defaultSize, page, counts.Ret.(int), defaultRenderPageItems, "/tags/"+strconv.FormatInt(id, 10)),
		"Tag":     tag,
		"Haohuos": haohuos.Ret,
		"Counts":  counts.Ret,
	})
}
