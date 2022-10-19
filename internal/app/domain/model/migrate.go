package model

import (
	"gorm.io/gorm"
)

// 注册数据库表
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		User{},
		Role{},
		Menu{},
		Resource{},
	)
}
