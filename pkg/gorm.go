package pkg

import (
	config "gin_template/config"
	"gorm.io/gorm"
)

func NewDB(config config.Application) *gorm.DB {
	switch config.System.DbType {
	case "mysql":
		return GormMysql(config.Mysql)
	default:
		return GormMysql(config.Mysql)
	}
}
