package orm

import "shahaohuo.com/shahaohuo/pkg/model"

type Link struct {
	model.Model
	UserId   string `gorm:"size:32;not null"sql:"index"`
	HaohuoId string `gorm:"size:32;not null"sql:"index"`
	Title    string `gorm:"size:1024"`
	Url      string `gorm:"text"`
	SiteName string `gorm:"size:12"sql:"index"`
}
