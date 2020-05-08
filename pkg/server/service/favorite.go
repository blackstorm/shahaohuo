package service

import (
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
)

func HaohuoFavorite(haohuoId, userId string) error {
	if f, e := orm.FindFavoriteByUserIdAndHaohuoId(userId, haohuoId); e != nil {
		logrus.Error(e)
		return ServiceError
	} else {
		// TODO check haohuo exist
		if f == nil {
			if e := orm.NewFavorite(haohuoId, userId).Create(); e != nil {
				logrus.Error(e)
				return ServiceError
			}
		}
		return nil
	}
}

func FindAllFavoritesByUserId(userId string) ([]orm.Favorite, error) {
	fs, err := orm.FinAllFavoritesByUserId(userId)
	if err != nil {
		logrus.Error(err)
		return nil, ServiceError
	}
	return fs, nil
}

func isUserFavoriteHaohuo(userId string, haohuoIds []string) ([]bool, error) {
	indexs := make(map[string]int)
	for i, v := range haohuoIds {
		indexs[v] = i
	}
	ret := make([]bool, len(haohuoIds))
	fs, err := orm.FindFavoritesByUserIdAndHaohuoIds(userId, haohuoIds)
	if err != nil {
		for _, f := range fs {
			ret[indexs[f.HaohuoId]] = true
		}
	} else {
		logrus.Error(err)
		return nil, ServiceError
	}
	return ret, nil
}
