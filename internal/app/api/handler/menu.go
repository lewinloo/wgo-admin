package handler

import (
	"gin_template/internal/app/common"
	"gin_template/internal/app/domain/service"
	"gin_template/internal/app/domain/service/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"strconv"
)

var MenuSet = wire.NewSet(wire.Struct(new(MenuHandler), "*"))

type MenuHandler struct {
	MenuSvc *service.MenuService
}

// @Tags 菜单模块
// @Summary 创建菜单接口
// @accept application/json
// @Security Bearer
// @Produce application/json
// @Param params body dto.CreateMenuParams true "创建菜单参数"
// @Success 200 {object} common.Result "成功返回体"
// @Router /menu [post]
func (h MenuHandler) Create(c *gin.Context) {
	var params dto.CreateMenuParams
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ResponseFail(c, 100001, "参数格式错误")
	} else {
		if success := h.MenuSvc.Create(params); success {
			common.ResponseOk(c, success)
		} else {
			common.ResponseFail(c, 100020, "创建失败")
		}
	}
}

// @Tags 菜单模块
// @Summary 查询菜单树
// @Security Bearer
// @accept application/json
// @Produce application/json
// @Success 200 {object} common.Result "成功返回体"
// @Router /menu/tree [get]
func (h MenuHandler) FindMenuTree(c *gin.Context) {
	list, err := h.MenuSvc.GetMenuTree()
	if err != nil {
		common.ResponseFail(c, 100000, "查询失败")
	} else {
		common.ResponseOk(c, list)
	}
}

// @Tags 菜单模块
// @Summary 删除菜单
// @Security Bearer
// @accept application/json
// @Produce application/json
// @Param id path int true "菜单id"
// @Success 200 {object} common.Result "成功返回体"
// @Router /menu/{id} [delete]
func (h MenuHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		common.ResponseFail(c, 100001, "请传递整型数字参数")
		return
	}
	err = h.MenuSvc.DeleteMenu(aid)
	if err != nil {
		common.ResponseFail(c, 100002, err.Error())
	} else {
		common.ResponseOk(c, nil)
	}
}
