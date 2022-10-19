package main

import (
	"flag"
	"gin_template/internal/app"
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

	app.LoadResource(configFilePath).Run()
}
