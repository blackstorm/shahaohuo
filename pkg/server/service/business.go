package service

import (
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/server/dto"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/storage"
	"time"
)

type businessHaohuoFunc func() ([]orm.BusinessHaohuo, error)
type basicUserFunc func() ([]orm.BasicUser, error)

func FindBasicUserById(id string) (*orm.BasicUser, error) {
	u, e := orm.FindBasicUserById(id)
	if e == nil {
		return u, nil
	}
	logrus.Error(e)
	return nil, ServiceError
}

func FindLatestUsersByLimit(limit int) ([]orm.BasicUser, error) {
	return basicUser(func() (user []orm.BasicUser, err error) {
		return orm.FindBasicUsersByOrderByAndLimit("created_at des", "desc", limit)
	})
}

func basicUser(fn basicUserFunc) ([]orm.BasicUser, error) {
	us, e := fn()
	if e == nil {
		return us, nil
	}
	logrus.Error(e)
	return nil, ServiceError
}

func FindUserBusinessHaohuosByLimit(size int) ([]orm.BusinessHaohuo, error) {
	return businessHaohuo(func() (haohuos []orm.BusinessHaohuo, err error) {
		return orm.FindBusinessHaohuosByLimit(size)
	})
}

func FindUserBusinessHaohuosById(id string) (*orm.BusinessHaohuo, error) {
	h, e := orm.FindBusinessHaohuosById(id)
	if e == nil {
		if h == nil {
			return nil, ResourcesNotFound
		}
		_setImageUrl(h)
		return h, nil
	}
	logrus.Error(e)
	return nil, ServiceError
}

func FindUserBusinessHaohuosByUserIdAndLimit(userId string, limit int) ([]orm.BusinessHaohuo, error) {
	return businessHaohuo(func() (haohuos []orm.BusinessHaohuo, err error) {
		return orm.FindBusinessHaohuosByUserIdAndLimit(userId, limit)
	})
}

func FindMostFavoriteBusinessHaohuosByDateAndLimit(start time.Time, end time.Time, limit int) ([]orm.BusinessHaohuo, error) {
	return businessHaohuo(func() (haohuos []orm.BusinessHaohuo, err error) {
		return orm.FindMostFavoriteBusinessHaohuosByDateAndLimit(start, end, limit)
	})
}

func Statics() *dto.Statics {
	counts, e := orm.CountUserAndHaohuo()
	if e != nil {
		logrus.Error(e)
	}
	return dto.NewStatics(counts)
}

func businessHaohuo(fn businessHaohuoFunc) ([]orm.BusinessHaohuo, error) {
	hs, e := fn()
	if e == nil {
		setImageUrl(hs)
		return hs, nil
	}
	logrus.Error(e)
	return nil, ServiceError
}

func setImageUrl(hs []orm.BusinessHaohuo) {
	for i, h := range hs {
		hs[i].ImageUrl = storage.ComplementImageUrl(&h)
	}
}

func _setImageUrl(h *orm.BusinessHaohuo) {
	h.ImageUrl = storage.ComplementImageUrl(h)
}
