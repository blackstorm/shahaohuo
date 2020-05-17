package orm

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/model"
	"time"
)

type Haohuo struct {
	model.Model
	Id          string `gorm:"primary_key; size:32"`
	UserId      string `gorm:"size:32;not null"sql:"index"`
	Name        string `gorm:"size:512;not null"sql:"index"`
	Price       int    `gorm:"type:int(10);not null"`
	Description string `gorm:"type:text;not null"`
	ImageUrl    string `gorm:"size:1024; not null"`
	Clicks      int    `gorm:"not null;default 0"`
}

type BusinessHaohuo struct {
	model.Model
	Id          string `json:"id"`
	UserName    string `json:"user_name"`
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Favorites   uint64 `json:"favorites"`
}

func (h *BusinessHaohuo) GetBaseImageUrl() string {
	return h.ImageUrl
}

func (h *BusinessHaohuo) SetFullImageUrl(url string) {
	h.ImageUrl = url
}

func (h *Haohuo) Create() error {
	return database.Create(h).Error
}

func (h *Haohuo) Update() error {
	return database.Save(h).Error
}

func (h *Haohuo) AsyncUpdateClicks() {
	go func() {
		// 解决方案
		rows := database.Model(h).Where("clicks = ?", h.Clicks).Update("clicks", h.Clicks+1).RowsAffected
		if rows <= 0 {
			logrus.Error("may update haohuo click fadile")
		}
	}()
}

func FindHaohuoById(id string) (*Haohuo, error) {
	haohuo := &Haohuo{}
	if e := database.Where("id = ?", id).First(haohuo).Error; e == nil {
		return haohuo, nil
	} else {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
}

func FindHaohuosByLimit(size int) (*[]Haohuo, error) {
	var hs []Haohuo
	if e := database.Order("created_at desc").Limit(size).Find(&hs).Error; e == nil {
		return &hs, nil
	} else {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
}

func CountHaohuosByUserId(userId string) int {
	var counts int
	database.Model(&Haohuo{}).Where("user_id = ?", userId).Count(&counts)
	return counts
}

const findBusinessHaohuosByIdsSQL = "SELECT h.*, COUNT(f.haohuo_id) as favorites FROM (SELECT h.*, u.name as user_name FROM haohuo as h force index(PRI), user as u WHERE h.id in (?) and u.id = h.user_id order by h.created_at desc) as h left join favorite as f ON f.haohuo_id = h.id group by h.id"

func FindBusinessHaohuosByIds(ids []string) ([]BusinessHaohuo, error) {
	var hs []BusinessHaohuo
	if e := database.Raw(findBusinessHaohuosByIdsSQL, ids).Scan(&hs).Error; e == nil {
		return hs, nil
	} else {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
}

func FindBusinessHaohuosByLimit(size int) ([]BusinessHaohuo, error) {
	var hs []BusinessHaohuo
	if e := database.Raw("SELECT h.*, COUNT(f.haohuo_id) as favorites FROM (SELECT h.*, u.name as user_name FROM haohuo as h, user as u WHERE u.id = h.user_id order by h.created_at desc limit ?) as h left join favorite as f ON f.haohuo_id = h.id group by h.id order by h.created_at desc", size).Scan(&hs).Error; e == nil {
		return hs, nil
	} else {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
}

func FindBusinessHaohuosById(id string) (*BusinessHaohuo, error) {
	var h BusinessHaohuo
	if e := database.Raw("SELECT h.*, COUNT(f.haohuo_id) as favorites FROM (SELECT h.*, u.name as user_name FROM haohuo as h, user as u WHERE u.id = h.user_id and h.id = ?) as h left join favorite as f ON f.haohuo_id = h.id group by h.id", id).Scan(&h).Error; e == nil {
		return &h, nil
	} else {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
}

func FindBusinessHaohuosByUserIdAndLimit(userId string, size int) ([]BusinessHaohuo, error) {
	var hs []BusinessHaohuo
	if e := database.Raw("SELECT h.*, COUNT(f.haohuo_id) as favorites FROM (SELECT h.*, u.name as user_name FROM haohuo as h, user as u WHERE h.user_id = ? and u.id = h.user_id order by h.created_at desc limit ?) as h left join favorite as f ON f.haohuo_id = h.id group by h.id", userId, size).Scan(&hs).Error; e == nil {
		return hs, nil
	} else {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
}

func FindBusinessHaohuosByUserIdAndPage(userId string, page, size int) ([]BusinessHaohuo, error) {
	var hs []BusinessHaohuo
	offset := (page - 1) * size
	ret := database.Raw("SELECT h.*, COUNT(f.haohuo_id) as favorites FROM (SELECT h.*, u.name as user_name FROM haohuo as h, user as u WHERE h.user_id = ? and u.id = h.user_id order by h.created_at desc limit ? offset ?) as h left join favorite as f ON f.haohuo_id = h.id group by h.id", userId, size, offset).Scan(&hs)
	if ret.RecordNotFound() {
		return nil, nil
	}
	if ret.Error != nil {
		return nil, ret.Error
	}
	return hs, nil
}

const findMostFavoriteBusinessHaohuosByDateAndLimitSQL = "SELECT f.count as favorites, h.*, u.name as user_name from (SELECT haohuo_id as id, COUNT(1) as count FROM favorite as f WHERE f.created_at between ? AND ? group by f.haohuo_id order by count desc limit ?) as f, haohuo as h, user as u where h.id = f.id and u.id = h.user_id"

func FindMostFavoriteBusinessHaohuosByDateAndLimit(start time.Time, end time.Time, size int) ([]BusinessHaohuo, error) {
	var hs []BusinessHaohuo
	if e := database.Raw(findMostFavoriteBusinessHaohuosByDateAndLimitSQL, start, end, size).Scan(&hs).Error; e == nil {
		return hs, nil
	} else {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
}
