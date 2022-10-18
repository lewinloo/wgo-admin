package handler

import (
	"gin_template/internal/app/common"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var HelloSet = wire.NewSet(wire.Struct(new(HelloHandler), "*"))

type HelloHandler struct {
}

// @Tags test
// @Summary 健康检查
// @accept application/json
// @Produce application/json
// @Router /hello [get]
func (h HelloHandler) Hello(c *gin.Context) {
	common.ResponseOk(c, nil)
}
