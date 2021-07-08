package model

import (
	"time"
)

// Article struct
type Article struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"user_id" gorm:"comment:用户ID"`
	Title     string    `json:"title" gorm:"comment:标题"`
	Content   string    `json:"content" gorm:"comment:用户登录密码;type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User *User
}
