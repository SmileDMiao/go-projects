package model

import (
	"time"
)

// User struct
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"userName" gorm:"comment:用户登录名;uniqueIndex"`
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg string    `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Articles []*Article
}
