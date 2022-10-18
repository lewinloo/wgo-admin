package handler

import (
	"gin_template/internal/app/common"
	"github.com/gin-gonic/gin"
)

func NewHello() HelloApi {
	return HelloApi{}
}

type HelloApi struct {
}

// @Tags test
// @Summary 健康检查
// @accept application/json
// @Produce application/json
// @Router /hello [get]
func (h HelloApi) Hello(c *gin.Context) {
	common.ResponseOk(c, nil)
}
