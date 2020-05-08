package model

import (
	"time"
)

type Model struct {
	CreatedAt time.Time `json:"created_at"sql:"index"`
	UpdatedAt time.Time `json:"updated_at"sql:"index"`
}

type SoftDeleteModel struct {
	CreatedAt time.Time  `json:"created_at"sql:"index"`
	UpdatedAt time.Time  `json:"updated_at"sql:"index"`
	DeletedAt *time.Time `sql:"index"`
}
