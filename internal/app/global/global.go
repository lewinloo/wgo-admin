package global

import (
	config "gin_template/config"

	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB           // 数据库
	CONFIG config.Application // 配置
	VP     *viper.Viper       // 读取配置文件库
	LOG    *zap.Logger        // 日志
	REDIS  *redis.Client      // redis
	CRON   *cron.Cron         // cron 定时任务
	CASBIN *casbin.Enforcer   // casbin
)
