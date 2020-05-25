package orm

import (
	"shahaohuo.com/shahaohuo/pkg/model"
)

type Video struct {
	model.Model
	UserId     string `gorm:"size:32;not null"sql:"index"`
	HaohuoId   string `gorm:"size:32;not null"sql:"index"`
	Source     string `gorm:"text;not null"`
	SourceType string `gorm:"size:128;not null"`
}

func (v *Video) IsBilibiliIframeVideo() bool {
	return v.SourceType == "BilibiliHTMLIframe"
}

func (v *Video) Create() error {
	return database.Create(v).Error
}

func NewBilibiliIframeVideo(userId, haohuoId, source string) *Video {
	return &Video{
		UserId:     userId,
		HaohuoId:   haohuoId,
		Source:     source,
		SourceType: "BilibiliHTMLIframe",
	}
}

func FindVideosByHaohuoId(id string) []Video {
	var videos []Video
	if database.Where("haohuo_id = ?", id).Find(&videos).RecordNotFound() {
		return nil
	}
	return videos
}
