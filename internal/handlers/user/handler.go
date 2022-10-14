package userHandlers

import (
	"gin_template/internal/common"
	"gin_template/internal/dto"
	userService "gin_template/internal/services/user"

	"github.com/gin-gonic/gin"
)

// @Tags 用户模块
// @Summary 用户注册接口
// @accept application/json
// @Produce application/json
// @Param params body dto.RegisterParams true "注册参数"
// @Success 200 {object} common.Result "成功返回体"
// @Router /api/v1/user/register [post]
func Register(c *gin.Context) {
	var params dto.RegisterParams
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ResponseFail(c, 100001, "参数格式错误")
	} else {
		if success, err := userService.Register(params); success {
			common.ResponseOk(c, success)
		} else {
			common.ResponseFail(c, 100010, err.Error())
		}
	}
}

// @Tags 用户模块
// @Summary 用户登录接口
// @accept application/json
// @Produce application/json
// @Param params body dto.LoginParams true "登录参数"
// @Success 200 {object} common.Result "成功返回体"
// @Router /api/v1/user/login [post]
func Login(c *gin.Context) {
	var params dto.LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ResponseFail(c, 100001, "参数格式错误")
	} else {
		if token, err := userService.Login(params); err != nil {
			common.ResponseFail(c, 100020, err.Error())
		} else {
			common.ResponseOk(c, gin.H{"token": token})
		}
	}
}
