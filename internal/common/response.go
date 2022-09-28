package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseOk(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Result{
		Success: true,
		Code:    200,
		Message: "本次请求成功",
		Data:    data,
	})
}

func ResponseFail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Result{
		Success: false,
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
