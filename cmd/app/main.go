package main

import (
	"flag"
	"gin_template/internal/app/core"
	"gin_template/internal/app/domain/model"
	"gin_template/internal/app/global"
	"gin_template/pkg"
	"go.uber.org/zap"
)

var (
	configFilePath string
)

// @title Gin Project API
// @basePath /api/v1
// @host 127.0.0.1:3000
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// 命令行参数
	flag.StringVar(&configFilePath, "f", "配置文件路径", "")
	flag.Parse()

	global.VP = pkg.Viper(configFilePath)
	global.LOG = pkg.Zap()
	zap.ReplaceGlobals(global.LOG)
	// 初始化数据库
	global.DB = pkg.NewDB(global.CONFIG)
	if global.DB != nil {
		// 初始化数据库表
		model.RegisterTables(global.DB)
		db, _ := global.DB.DB()
		defer db.Close()
	}

	core.RunServer()
}
