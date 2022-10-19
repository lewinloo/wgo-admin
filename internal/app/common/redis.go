package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

// 使用 redis 模块管理一个 key
type RedisModule struct {
	Client *redis.Client
	Key    string
}

func NewRedisModule(client *redis.Client, key string) *RedisModule {
	return &RedisModule{
		Client: client,
		Key:    key,
	}
}

func (r *RedisModule) SetRedisKey(value any, ex time.Duration) error {
	return r.Client.Set(ctx, r.Key, value, ex).Err()
}

func (r *RedisModule) GetRedisKey() (result any, err error) {
	val, err := r.Client.Get(ctx, r.Key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("Redis键名为%s的值为空", r.Key)
	} else if err != nil {
		return nil, err
	} else {
		return val, nil
	}
}

func (r *RedisModule) DelRedisKey() error {
	return r.Client.Del(ctx, r.Key).Err()
}

func (r *RedisModule) CloseRedis() {
	err := r.Client.Close()
	if err != nil {
		fmt.Println("close redis error")
	}
}
