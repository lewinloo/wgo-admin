package handler

import (
	"gin_template/internal/app/common"
	"gin_template/internal/app/domain/service"
	"gin_template/internal/app/domain/service/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserHandler), "*"))

type UserHandler struct {
	UserSvc *service.UserService
}

// @Tags 用户模块
// @Summary 用户注册接口
// @Security Bearer
// @accept application/json
// @Produce application/json
// @Param params body dto.RegisterParams true "注册参数"
// @Success 200 {object} common.Result "成功返回体"
// @Router /user/register [post]
func (h UserHandler) Register(c *gin.Context) {
	var params dto.RegisterParams
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ResponseFail(c, 100001, "参数格式错误")
	} else {
		if success, err := h.UserSvc.Register(params); success {
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
// @Router /user/login [post]
func (h UserHandler) Login(c *gin.Context) {
	var params dto.LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ResponseFail(c, 100001, "参数格式错误")
	} else {
		if token, err := h.UserSvc.Login(params); err != nil {
			common.ResponseFail(c, 100020, err.Error())
		} else {
			common.ResponseOk(c, gin.H{"token": token})
		}
	}
}

// @Tags 用户模块
// @Summary 用户列表
// @Security Bearer
// @accept application/json
// @Produce application/json
// @Param params body dto.QueryUserListParams true "登录参数"
// @Success 200 {object} common.Result "成功返回体"
// @Router /user/list [post]
func (h UserHandler) List(c *gin.Context) {
	var params dto.QueryUserListParams
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ResponseFail(c, 100001, "参数格式错误")
	} else {
		if res, err := h.UserSvc.GetList(params); err != nil {
			common.ResponseFail(c, 400, err.Error())
		} else {
			common.ResponseOk(c, res)
		}
	}
}

// @Tags 用户模块
// @Summary 删除用户
// @Security Bearer
// @accept application/json
// @Produce application/json
// @Param params body dto.Ids true "删除用户参数"
// @Success 200 {object} common.Result "成功返回体"
// @Router /user/delete [delete]
func (h UserHandler) Delete(c *gin.Context) {
	var params dto.Ids
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ResponseFail(c, 100001, "参数格式错误")
	} else {
		if success := h.UserSvc.DeleteUsers(params); !success {
			common.ResponseFail(c, 400, "删除失败")
		} else {
			common.ResponseOk(c, nil)
		}
	}
}
