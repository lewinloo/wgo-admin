package main

import (
	"gin_template/internal/core"
	"gin_template/internal/global"
	"gin_template/pkg"

	"go.uber.org/zap"
)

// @title Gin Project API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host 127.0.0.1:3000
// @BasePath /api/v1
func main() {
	global.VP = pkg.Viper()
	global.LOG = pkg.Zap()
	zap.ReplaceGlobals(global.LOG)
	// 初始化数据库
	global.DB = pkg.NewDB()
	if global.DB != nil {
		// 初始化数据库表
		pkg.RegisterTables(global.DB)
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
