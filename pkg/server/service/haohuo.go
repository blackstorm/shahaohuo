package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/server/dto"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"time"
)

// TODO tx support
func CreateOrUpdate(req *dto.HaohuoRequest, userId string) (*orm.Haohuo, error) {
	if h, e := orm.FindHaohuoById(req.Id); e != nil {
		logrus.Error(e)
		return nil, ServiceError
	} else {
		if h == nil {
			// do save
			h = &orm.Haohuo{}
			_ = copier.Copy(h, req)
			h.UserId = userId
			e = h.Create()
		} else {
			// do upodate
			if h.UserId != userId {
				return nil, ResourcesExist
			}
			h.UpdatedAt = time.Now()
			h.Name = req.Name
			h.Price = req.Price
			h.Description = req.Description
			h.ImageUrl = req.ImageUrl
			e = h.Update()
		}

		if e != nil {
			logrus.Error(e)
			return nil, ServiceError
		}

		// save or update tags
		e = SaveOrUpdateHaohuoTags(h.Id, req.Tags)
		if e != nil {
			logrus.Error(e)
			return nil, ServiceError
		}

		return h, nil
	}
}

func FindHaohuoById(id string) (*orm.Haohuo, error) {
	h, e := orm.FindHaohuoById(id)
	if e != nil {
		logrus.Error(e)
		return nil, ServiceError
	}
	return h, nil
}
