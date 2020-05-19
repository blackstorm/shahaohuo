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

func FindByUserIdAndPath(userId, path string) (*Image, error) {
	var image Image
	_db := database.Where("user_id = ? and path = ?", userId, path).First(&image)
	if _db.RecordNotFound() {
		return nil, nil
	}
	return &image, _db.Error
}
