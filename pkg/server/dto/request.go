package dto

type CommentRequest struct {
	Content string `json:"content"`
}

type PatchUserRequest struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"binding:"max=8,min=1"`
	Bio    string `json:"bio"binding:"max=128"`
}

type HaohuoVideoRequest struct {
	Url string `json:"url"`
}
