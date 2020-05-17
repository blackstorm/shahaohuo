package service

import (
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/crypto"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/storage"
)

const (
	defaultAvatarPath = "/default/avatar.png"
)

func Login(id string, passwd string) (*orm.User, error) {
	if u, e := orm.FindUserById(id); e != nil {
		logrus.Error(e)
		return nil, e
	} else {
		if u != nil {
			if err := crypto.Compare(u.Password, passwd); err == nil {
				return u, nil
			}
		}
		return nil, UsernameOrPasswordError
	}
}

func Register(id string, passwd string) (*orm.User, error) {
	if u, e := orm.FindUserById(id); e == nil {
		if u != nil {
			return nil, ResourcesExist
		}
		// encode password
		var encoded string
		if encoded, e = crypto.Encrypt(passwd); e == nil {
			user := &orm.User{
				Id:       id,
				Name:     id,
				Password: encoded,
				Avatar:   defaultAvatarPath,
			}
			if e = user.Save(); e == nil {
				return user, nil
			} else {
				logrus.Error(e)
			}
		} else {
			logrus.Error(e)
		}
	} else {
		logrus.Error(e)
	}

	return nil, ServiceError
}

func FindUserById(id string) (*orm.User, error) {
	if u, e := orm.FindUserById(id); e == nil {

		storage.AutoComplementImageUrl(u)
		return u, nil
	} else {
		logrus.Error(e)
		return nil, ServiceError
	}
}
