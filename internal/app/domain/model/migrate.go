package model

import (
	"gin_template/internal/app/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

// 注册数据库表
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		User{},
		Role{},
		Menu{},
	)

	if err != nil {
		global.LOG.Error("注册数据库表失败", zap.Error(err))
		os.Exit(0)
	}

	global.LOG.Info("注册数据库表成功！！！")
}
