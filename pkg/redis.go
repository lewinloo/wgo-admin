package pkg

import (
	"gin_template/internal/app/config"
	"github.com/go-redis/redis/v8"
)

func NewRedis(redisCfg config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	return client
}
