package orm

import (
	"github.com/jinzhu/gorm"
	"shahaohuo.com/shahaohuo/pkg/model"
)

type Favorite struct {
	model.Model
	UserId   string `gorm:"size:32;not null;unique_index:user_haohuo_index"sql:"index"`
	HaohuoId string `gorm:"size:32;not null;unique_index:user_haohuo_index"sql:"index"`
}

type HaohuoFavorites struct {
	HaohuoId  string `json:"haohuo_id"`
	Favorites uint64 `json:"favorites"`
}

type FavoriteUser struct {
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	HaohuoId string `json:"haohuo_id"`
}

func (f *Favorite) Create() error {
	return database.Create(f).Error
}

func NewFavorite(haohuoId, userId string) *Favorite {
	return &Favorite{
		UserId:   userId,
		HaohuoId: haohuoId,
	}
}

func CountFavoritesByHaohuoId(haohuoId string) (*HaohuoFavorites, error) {
	var hf HaohuoFavorites
	e := database.Raw("select haohuo_id, count(1) as favorites from favorite where haohuo_id = ? group by haohuo_id", haohuoId).Scan(&hf).Error
	if e != nil {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
	return &hf, nil
}

func FindFavoriteByUserIdAndHaohuoId(userId, haohuoId string) (*Favorite, error) {
	var f Favorite
	if e := database.Where("user_id = ? and haohuo_id = ?", userId, haohuoId).First(&f).Error; e != nil {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
	return &f, nil
}

func FinAllFavoritesByUserId(userId string) ([]Favorite, error) {
	var fs []Favorite
	if e := database.Where("user_id = ?", userId).Find(fs).Error; e != nil {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
	return fs, nil
}

func FindFavoritesByUserIdAndHaohuoIds(userId string, haohuoIds []string) ([]Favorite, error) {
	var fs []Favorite
	if e := database.Where("user_id = ? and haohuo_id in (?)", userId, haohuoIds).Find(fs).Error; e != nil {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
	return fs, nil
}

const findFavoriteUsersByHaohuoIdSQL = "select f.haohuo_id, u.id as user_id, u.name as user_name from favorite as f, user as u where u.id = f.user_id and f.haohuo_id = ? order by f.created_at desc limit ?"
func FindFavoriteUsersByHaohuoId(id string, limit uint) ([]FavoriteUser, error) {
	var us []FavoriteUser
	if e := database.Raw(findFavoriteUsersByHaohuoIdSQL, id, limit).Scan(&us).Error; e != nil {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
	return us, nil
}
