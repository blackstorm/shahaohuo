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
	HaohuoId string `gorm:"size:32;not null;unique_index:tag_haohuo_id_index"sql:"index"`
	TagId    int64  `gorm:"not null;unique_index:tag_haohuo_id_index"sql:"index"`
}

func (h *HaohuoTag) Create() error {
	return database.Create(h).Error
}

func FindAllTag() ([]Tag, error) {
	var tags []Tag
	e := database.Find(&tags).Error
	return tags, e
}

type HaohuoIds struct {
	HaohuoId []string
}

func FindHaohuoIdsByTagId(tagId int64, page int, size int) []string {
	var htags []HaohuoTag
	database.Where("tag_id = ?", tagId).Order("created_at DESC").Offset((page - 1) * size).Limit(size).Find(&htags)
	if len(htags) > 0 {
		ids := make([]string, len(htags))
		for i, t := range htags {
			ids[i] = t.HaohuoId
		}
		return ids
	}
	return nil
}

func FindTagsByHaohuoId(id string) []Tag {
	var tags []Tag
	database.Raw("select t.* from haohuo_tag as ht, tag as t where t.id = ht.tag_id and ht.haohuo_id = ?", id).Scan(&tags)
	return tags
}

func CountHaohuosByTagId(id int64) int {
	var counts int
	database.Model(&HaohuoTag{}).Where("tag_id = ?", id).Count(&counts)
	return counts
}

func FindTagById(id int64) *Tag {
	var tag Tag
	if database.First(&tag, id).RecordNotFound() {
		return nil
	}
	return &tag
}

func DeleteHaohuoTagByHaohuoIdAndTagId(haohuoId string, tagId int64) error {
	return database.Where("haohuo_id = ? and tagId = ?", haohuoId, tagId).Delete(&HaohuoTag{}).Error
}
