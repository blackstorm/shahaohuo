package orm

import (
	"shahaohuo.com/shahaohuo/pkg/model"
	"time"
)

// TODO
type Comment struct {
	model.SoftDeleteModel
	Id       string `gorm:"primary_key; size:32"`
	UserId   string `gorm:"size:32;not null"sql:"index"`
	HaohuoId string `gorm:"size:32;not null"sql:"index"`
	Content  string `gorm:"type:text;not null"`
}

type HaohuoComment struct {
	Id        string    `json:"id"`
	UserName  string    `json:"user_name"`
	UserId    string    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserComment struct {
	Id         string    `json:"id"`
	HaohuoId   string    `json:"haohuo_id"`
	HaohuoName string    `json:"haohuo_name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (c *Comment) Create() error {
	return database.Save(c).Error
}

func (c *Comment) Update() error {
	database.Where("id = ?", c.Id).Updates(c, true)
	return database.Save(c).Error
}

const findUserCommentsByHaohuoIdSQL = "select u.id as user_id, u.name as user_name, c.* from comment as c, user as u where u.id = c.user_id and c.haohuo_id = ? order by c.created_at desc limit ?"

func FindHaohuoCommentsByHaohuoId(id string, limit uint) ([]HaohuoComment, error) {
	var ucs []HaohuoComment
	e := database.Raw(findUserCommentsByHaohuoIdSQL, id, limit).Scan(&ucs).Error
	return ucs, e
}

const findUserCommentsByUserIdSQL = "select h.id as haouser_id, h.name as haohuo_name, c.* from haohuo as h, comment as c where h.id = c.haohuo_id and c.user_id = ? order by created_at desc limit ?"

func FindUserCommentsByUserId(id string, limit uint) ([]UserComment, error) {
	var ucs []UserComment
	e := database.Raw(findUserCommentsByUserIdSQL, id, limit).Scan(&ucs).Error
	return ucs, e
}

func DeleteUserCommentById(id string) error {
	return database.Where("id = ?", id).Delete(&HaohuoComment{}).Error
}
