package orm

import (
	"shahaohuo.com/shahaohuo/pkg/model"
)

type Tag struct {
	model.Model
	Id   int64  `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"varchar(255);not null;unique_index:tag_name_index"sql:"index"`
}

func (t Tag) SeoKeyWorld() string {
	return t.Name
}

type HaohuoTag struct {
	model.Model
	Id       int64  `gorm:"primary_key;AUTO_INCREMENT"`
	HaohuoId string `gorm:"size:32;not null"`
	TagId    int64  `gorm:"not null"`
}

func (h *HaohuoTag) Create() error {
	return database.Create(h).Error
}

func FindAllTag() ([]Tag, error) {
	var tags []Tag
	e := database.Find(&tags).Error
	return tags, e
}

func FindTagsByHaohuoId(id string) []Tag {
	var tags []Tag
	database.Raw("select t.* from haohuo_tag as ht, tag as t where t.id = ht.tag_id and ht.haohuo_id = ?", id).Scan(&tags)
	return tags
}

func DeleteHaohuoTagByHaohuoIdAndTagId(haohuoId string, tagId int64) error {
	return database.Where("haohuo_id = ? and tagId = ?", haohuoId, tagId).Delete(&HaohuoTag{}).Error
}
