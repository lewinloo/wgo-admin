package handler

import (
	"gin_template/internal/app/common"
	"gin_template/internal/app/domain/service"
	"gin_template/internal/app/domain/service/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"strconv"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleHandler), "*"))

type RoleHandler struct {
	RoleSvc *service.RoleService
}

// @Tags 角色模块
// @Summary 创建角色接口
// @Security Bearer
// @accept application/json
// @Produce application/json
// @Param params body dto.CreateRoleParams true "创建角色参数"
// @Success 200 {object} common.Result "成功返回体"
// @Router /role [post]
func (h RoleHandler) Create(c *gin.Context) {
	var params dto.CreateRoleParams
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ResponseFail(c, 100001, "参数格式错误")
	} else {
		if success := h.RoleSvc.Create(params); success {
			common.ResponseOk(c, success)
		} else {
			common.ResponseFail(c, 100010, err.Error())
		}
	}
}

// @Tags 角色模块
// @Summary 角色列表接口
// @Security Bearer
// @accept application/json
// @Produce application/json
// @Param params body dto.QueryRoleListParams true "查询列表参数"
// @Success 200 {object} common.Result "成功返回体"
// @Router /role/list [post]
func (h RoleHandler) FindList(c *gin.Context) {
	var params dto.QueryRoleListParams
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ResponseFail(c, 100001, "参数格式错误")
	} else {
		if result, err := h.RoleSvc.FindList(params); err == nil {
			common.ResponseOk(c, result)
		} else {
			common.ResponseFail(c, 400, err.Error())
		}
	}
}

// @Tags 角色模块
// @Summary 删除角色
// @Security Bearer
// @accept application/json
// @Produce application/json
// @Param id path int true "角色id"
// @Success 200 {object} common.Result "成功返回体"
// @Router /role/{id} [delete]
func (h RoleHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		common.ResponseFail(c, 100001, "请传递整型数字参数")
		return
	}
	err = h.RoleSvc.DeleteById(aid)
	if err != nil {
		common.ResponseFail(c, 100002, err.Error())
	} else {
		common.ResponseOk(c, nil)
	}
}
