package handler

import (
	"gin_template/internal/app/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var HelloSet = wire.NewSet(wire.Struct(new(HelloHandler), "*"))

type HelloHandler struct {
	Redis *redis.Client
	Log   *zap.Logger
}

// @Tags test
// @Summary 健康检查
// @accept application/json
// @Produce application/json
// @Router /hello [get]
func (h HelloHandler) Hello(c *gin.Context) {
	redisModule := common.NewRedisModule(h.Redis, "user:1")
	_ = redisModule.SetRedisKey("Wire", 0)
	h.Log.Info("redis set user:1 to Wire")
	common.ResponseOk(c, nil)
}
