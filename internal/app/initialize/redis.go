package initialize

import (
	"gin_template/internal/app/config"
	"gin_template/pkg"
	"github.com/go-redis/redis/v8"
)

func Redis() (*redis.Client, func(), error) {
	client := pkg.NewRedis(config.C.Redis)
	cleanFunc := func() {}
	return client, cleanFunc, nil
}
