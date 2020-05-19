package html

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/service"
	"shahaohuo.com/shahaohuo/pkg/util"
	"time"
)

var shanghai, _ = time.LoadLocation("Asia/Shanghai")

func Index(ctx *gin.Context) {

	hsChan := make(chan util.AsyncResult)
	defer close(hsChan)
	go func() {
		hs, e := service.FindUserBusinessHaohuosByLimit(12)
		hsChan <- util.AsyncResult{
			Ret:   hs,
			Error: e,
		}
	}()

	mostFavoriteChan := make(chan util.AsyncResult)
	defer close(mostFavoriteChan)
	go func() {
		mostFavorite, e := queryMostFavorite()
		mostFavoriteChan <- util.AsyncResult{
			Ret:   mostFavorite,
			Error: e,
		}
	}()

	latestUsersChan := make(chan util.AsyncResult)
	defer close(latestUsersChan)
	go func() {
		us, e := service.FindLatestUsersByLimit(64)
		latestUsersChan <- util.AsyncResult{
			Ret:   us,
			Error: e,
		}
	}()

	staticsChan := make(chan util.AsyncResult)
	defer close(staticsChan)
	go func() {
		ret := service.Statics()
		staticsChan <- util.AsyncResult{
			Ret:   ret,
			Error: nil,
		}
	}()

	tagsChan := make(chan util.AsyncResult)
	defer close(tagsChan)
	go func() {
		tags, e := orm.FindAllTag()
		tagsChan <- util.AsyncResult{
			Ret:   tags,
			Error: e,
		}
	}()

	hs := <-hsChan
	mostFavorite := <-mostFavoriteChan
	latestUsers := <-latestUsersChan
	statics := <-staticsChan
	tags := <-tagsChan

	if util.CheckAsyncResultsError(hs, mostFavorite, tags, statics, latestUsers) {
		errorPage(ctx)
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"NewHaohuos":   hs.Ret,
		"MostFavorite": mostFavorite.Ret,
		"LatestUsers":  latestUsers.Ret,
		"Statics":      statics.Ret,
		"Tags":         tags.Ret,
	})

}

func queryMostFavorite() ([]orm.BusinessHaohuo, error) {
	now := time.Now().In(shanghai)
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return service.FindMostFavoriteBusinessHaohuosByDateAndLimit(start, now, 4)
}
