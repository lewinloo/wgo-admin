package helloHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags test
// @Summary 健康检查
// @accept application/json
// @Produce application/json
// @Router /api/v1/hello [get]
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
