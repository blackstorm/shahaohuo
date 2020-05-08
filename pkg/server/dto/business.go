package dto

import "shahaohuo.com/shahaohuo/pkg/server/orm"

type UserBusinessHaohuo struct {
	orm.BusinessHaohuo
	IsFavorite bool `json:"is_favorite"`
}

type HaohuoRequest struct {
	Id          string
	Name        string  `json:"name" binding:"required,max=64,min=1"`
	Price       int     `json:"price" binding:"required,gte=1,lte=99999"`
	Description string  `json:"description" binding:"required,max=2048,min=1"`
	ImageUrl    string  `json:"imageUrl" binding:"required"`
	Tags        []int64 `json:"tags";binding:"required,min=1,max=5"`
}

type Statics struct {
	Counts *orm.UserAndHaohuoCount
}

func NewStatics(counts *orm.UserAndHaohuoCount) *Statics {
	return &Statics{Counts: counts}
}
