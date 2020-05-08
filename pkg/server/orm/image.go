package orm

import (
	"shahaohuo.com/shahaohuo/pkg/model"
)

type Image struct {
	model.Model
	Id     string `gorm:"primary_key; size:32"`
	UserId string `gorm:"size:32;not null"sql:"index"`
	Path   string `gorm:"not null;size:256"`
}

func SaveImage(userId, id, path string) error {
	return database.Create(&Image{
		UserId: userId,
		Id:     id,
		Path:   path,
	}).Error
}
