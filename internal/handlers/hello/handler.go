package helloHandlers

import (
	"gin_template/internal/common"

	"github.com/gin-gonic/gin"
)

// @Tags test
// @Summary 健康检查
// @accept application/json
// @Produce application/json
// @Router /api/v1/hello [get]
func Hello(c *gin.Context) {
	common.ResponseOk(c, nil)
}
