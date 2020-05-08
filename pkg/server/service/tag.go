package service

import (
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
)

func SaveOrUpdateHaohuoTags(haohuoId string, tagIds []int64) error {
	var e error
	tags := orm.FindTagsByHaohuoId(haohuoId)
	if tags != nil && len(tags) > 0 {
		tagMap := createTagMarkMap(tags)
		// mark
		for _, id := range tagIds {
			if _, ok := tagMap[id]; ok {
				tagMap[id] = true
			} else {
				// insert
				e = saveHaohuoTag(haohuoId, id)
			}
		}
		if e == nil {
			// delete not mark
			e = deleteHoahuoTagByTagMarkMap(haohuoId, tagMap)
		}
	} else {
		e = saveHaohuoTags(haohuoId, tagIds)
	}
	if e != nil {
		logrus.Error(e)
		return ServiceError
	}
	return nil
}

func createTagMarkMap(tags []orm.Tag) map[int64]bool {
	tagMap := make(map[int64]bool)
	// flag
	for _, tag := range tags {
		tagMap[tag.Id] = false
	}
	return tagMap
}

func deleteHoahuoTagByTagMarkMap(haohuoId string, markMap map[int64]bool) error {
	for k, v := range markMap {
		if !v {
			e := DeleteHaohuoTagByHaohuoIdAndTagId(haohuoId, k)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

func DeleteHaohuoTagByHaohuoIdAndTagId(haohuoId string, tagId int64) error {
	return orm.DeleteHaohuoTagByHaohuoIdAndTagId(haohuoId, tagId)
}

func saveHaohuoTag(haohuoId string, tagId int64) error {
	ht := &orm.HaohuoTag{
		HaohuoId: haohuoId,
		TagId:    tagId,
	}
	return ht.Create()
}

func saveHaohuoTags(haohuoId string, tagIds []int64) error {
	for _, id := range tagIds {
		e := saveHaohuoTag(haohuoId, id)
		if e != nil {
			return e
		}
	}
	return nil
}
