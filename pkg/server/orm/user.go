package orm

import (
	"github.com/jinzhu/gorm"
	"shahaohuo.com/shahaohuo/pkg/model"
)

type User struct {
	model.Model
	Id       string `gorm:"primary_key;size:32"`
	Name     string `gorm:"size:12;not null"`
	Password string `gorm:"size:64;not null"`
	Bio      string `gorm:"size:180"`
	Avatar   string `gorm:"size:1024; not null"`
}

type BasicUser struct {
	model.Model
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Save() error {
	if e := database.Create(u).Error; e != nil {
		return e
	}
	return nil
}

// TODO version
func (u *User) Update() error {
	return database.Save(u).Error
}

func (u *User) GetBaseImageUrl() string {
	return u.Avatar
}

func FindUserById(id string) (*User, error) {
	user := &User{}
	if err := database.Where("id = ?", id).First(user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return user, nil
	}
}

func FindBasicUsersByOrderByAndLimit(orderBy string, sort string, limit int) ([]BasicUser, error) {
	var us []BasicUser
	if e := database.Raw("select id, name, created_at, updated_at from user order by ? limit ?", orderBy+" "+sort, limit).Scan(&us).Error; e != nil {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
	return us, nil
}

func FindBasicUserById(id string) (*BasicUser, error) {
	var u BasicUser
	if e := database.Raw("select id, name, created_at, updated_at from user where id = ?", id).Scan(&u).Error; e != nil {
		if gorm.IsRecordNotFoundError(e) {
			return nil, nil
		}
		return nil, e
	}
	return &u, nil
}
