package html

import (
	"github.com/gin-gonic/gin"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/service"
	"shahaohuo.com/shahaohuo/pkg/server/storage"
	"shahaohuo.com/shahaohuo/pkg/server/web"
	"shahaohuo.com/shahaohuo/pkg/util"
)

/**
TODO
1。分类，标签等信息
2. 全部好货
3. 搜索框
*/
func Haohuos(c *gin.Context) {

}

func Haohuo(c *gin.Context) {
	var uriParams web.UriIdParams
	if e := c.BindUri(&uriParams); e != nil {
		// TODO bad page
		errorPage(c)
		return
	}

	id := uriParams.Id

	// find haohuo
	h, e := service.FindUserBusinessHaohuosById(id)
	if e != nil {
		if e == service.ResourcesNotFound {
			notFoundPage(c)
		} else {
			errorPage(c)
		}
		return
	}

	// find user
	userChannel := make(chan util.AsyncResult)
	defer close(userChannel)
	go func() {
		u, e := service.FindUserById(h.UserId)
		userChannel <- util.AsyncResult{
			Ret:   u,
			Error: e,
		}
	}()

	// favorites
	fusChannel := make(chan util.AsyncResult)
	defer close(fusChannel)
	go func() {
		fus, e := orm.FindFavoriteUsersByHaohuoId(id, 100)
		fusChannel <- util.AsyncResult{
			Ret:   fus,
			Error: e,
		}
	}()

	// comments
	cmsChannel := make(chan util.AsyncResult)
	defer close(cmsChannel)
	go func() {
		cms, e := orm.FindHaohuoCommentsByHaohuoId(id, 99)
		if len(cms) > 0 {
			for i, _ := range cms {
				storage.AutoComplementImageUrl(&cms[i])
			}
		}
		cmsChannel <- util.AsyncResult{
			Ret:   cms,
			Error: e,
		}
	}()

	// tags
	tgsChannel := make(chan util.AsyncResult)
	defer close(tgsChannel)
	go func() {
		tgs := orm.FindTagsByHaohuoId(id)
		tgsChannel <- util.AsyncResult{
			Ret:   tgs,
			Error: e,
		}
	}()

	u := <-userChannel
	fus := <-fusChannel
	cms := <-cmsChannel
	tgs := <-tgsChannel

	if util.CheckAsyncResultsError(u, fus, cms, tgs) {
		errorPage(c)
		return
	}

	c.HTML(200, "haohuo.html", gin.H{
		"Haohuo":        h,
		"User":          u.Ret,
		"FavoriteUsers": fus.Ret,
		"Comments":      cms.Ret,
		"Tags":          tgs.Ret,
	})
}
