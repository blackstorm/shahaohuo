package dto

import "time"

type HaohuoFavorites struct {
	HaohuoId  string `json:"haohuo_id"`
	Favorites uint64 `json:"favorites"`
}

type CommentResponse struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	HaohuoId  string    `json:"haohuo_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
