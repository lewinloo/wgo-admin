package core

import (
	"fmt"
	"gin_template/internal/global"
	"gin_template/internal/router"
	"gin_template/pkg"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	// 配置文件如果使用 redis 则初始化 redis 服务
	if global.CONFIG.System.UseRedis {
		pkg.Redis()
	}

	// 根据配置文件是否使用定时任务
	if global.CONFIG.System.UseCron {
		global.CRON = pkg.NewCron()
	}

	Router := router.New()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Port)
	s := initServer(address, Router)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("port", address))

	fmt.Printf(`
	Swagger文档地址:http://127.0.0.1%s/swagger/index.html
	服务地址:http://127.0.0.1%s%s
	`, address, address, global.CONFIG.System.GlobalPrefix)
	global.LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
