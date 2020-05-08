package service

import (
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/server/dto"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
)

func CommentHaohuo(haohuoId string, userId string, request dto.CommentRequest) (*orm.Comment, error) {
	h, e := orm.FindHaohuoById(haohuoId)
	if e == nil {
		if h == nil {
			return nil, ResourcesNotFound
		}
		comment := &orm.Comment{
			Id:       xid.New().String(),
			UserId:   userId,
			HaohuoId: h.Id,
			Content:  request.Content,
		}
		e := comment.Create()
		if e == nil {
			return comment, nil
		}
	}
	logrus.Error(e)
	return nil, e
}
