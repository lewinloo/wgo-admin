package redisService

import (
	"context"
	"fmt"
	"gin_template/internal/global"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// 使用 redis 模块管理一个 key
type redisModule struct {
	redisClient *redis.Conn
	key         string
}

func NewRedisModule(key string) *redisModule {
	return &redisModule{
		redisClient: global.REDIS.Conn(ctx),
		key:         key,
	}
}

var ctx = context.Background()

func (r *redisModule) SetRedisKey(value any, ex time.Duration) error {
	return r.redisClient.Set(ctx, r.key, value, ex).Err()
}

func (r *redisModule) GetRedisKey() (result any, err error) {
	val, err := r.redisClient.Get(ctx, r.key).Result()
	if err == redis.Nil {
		global.LOG.Info("获取 redis 值为空", zap.String("key", r.key))
		return nil, fmt.Errorf("Redis键名为%s的值为空", r.key)
	} else if err != nil {
		global.LOG.Error("获取 redis 值报错", zap.String("error", err.Error()))
		return nil, err
	} else {
		return val, nil
	}
}

func (r *redisModule) DelRedisKey() error {
	return r.redisClient.Del(ctx, r.key).Err()
}

func (r *redisModule) CloseRedis() {
	err := r.redisClient.Close()
	if err != nil {
		global.LOG.Error("关闭 Redis 连接失败", zap.String("error", err.Error()))
	}
}
