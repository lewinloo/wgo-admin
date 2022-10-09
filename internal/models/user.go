package models

import (
	"gin_template/internal/global"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	global.BaseModel
	Username     string  `gorm:"type:varchar(20);not null;unique" json:"username"`
	Password     string  `gorm:"size:255;not null" json:"password"`
	Mobile       string  `gorm:"type:varchar(11);not null;unique" json:"mobile"`
	Email        string  `gorm:"type:varchar(255);not null;unique" json:"email"`
	Avatar       string  `gorm:"type:varchar(255)" json:"avatar"`
	Nickname     string  `gorm:"type:varchar(20)" json:"nickname"`
	Introduction string  `gorm:"type:varchar(255)" json:"introduction"`
	Status       uint    `gorm:"type:tinyint(1);default:1;comment:'1正常, 2禁用'" json:"status"`
	IsAdmin      bool    `gorm:"default:0;comment:'0不是 1是'" json:"isAdmin"`
	Roles        []*Role `gorm:"many2many:user_roles" json:"roles"`
}

// 加密
func (user *User) SetPssword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// 验证密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
