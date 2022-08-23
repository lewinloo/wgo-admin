package pkg

import (
	"gin_template/internal/global"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// 注册数据库表
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate()

	if err != nil {
		global.LOG.Error("注册数据库表失败", zap.Error(err))
		os.Exit(0)
	}

	global.LOG.Info("注册数据库表成功！！！")
}
