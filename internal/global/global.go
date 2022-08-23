package global

import (
	config "gin_template/configs"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB      // 数据库
	CONFIG config.Server // 配置
	VP     *viper.Viper  // 读取配置文件库
	LOG    *zap.Logger   // 日志
	REDIS  *redis.Client // redis
)
