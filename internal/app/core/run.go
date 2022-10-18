package core

import (
	"fmt"
	"gin_template/internal/app/api"
	"gin_template/internal/app/global"
	"gin_template/pkg"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	// 设置生产模式
	if global.CONFIG.System.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 配置文件如果使用 redis 则初始化 redis 服务
	if global.CONFIG.System.UseRedis {
		global.REDIS = pkg.NewRedis(global.CONFIG.Redis)
	}

	// 根据配置文件是否使用定时任务
	if global.CONFIG.System.UseCron {
		global.CRON = pkg.NewCron()
	}

	// 根据配置文件是否使用casbin
	if global.CONFIG.System.UseCasbin {
		enforcer, err := pkg.NewCasbinEnforcer(global.DB)
		if err != nil {
			global.LOG.Error(fmt.Sprintf("初始化Casbin失败：%v", err))
			panic(any(fmt.Sprintf("初始化Casbin失败：%v", err)))
		}
		global.LOG.Info("初始化Casbin成功！")
		global.CASBIN = enforcer
	}

	Router := api.New()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Port)

	// 根据平台不同
	var s server
	systemType := runtime.GOOS
	if systemType == "windows" {
		s = initServerWin(address, Router)
	} else {
		s = initServerOther(address, Router)
	}

	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("port", address))

	fmt.Printf(`
	Swagger文档地址:http://127.0.0.1%s/swagger/index.html
	服务地址:http://127.0.0.1%s%s
	`, address, address, global.CONFIG.System.GlobalPrefix)
	global.LOG.Error(s.ListenAndServe().Error())
}
