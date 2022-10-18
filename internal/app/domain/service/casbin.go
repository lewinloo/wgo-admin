package service

import (
	"errors"
	"gin_template/internal/app/domain/service/dto"
	"gin_template/internal/app/global"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/google/wire"
)

var CasbinSet = wire.NewSet(wire.Struct(new(CasbinService), "*"))

type CasbinService struct{}

// 更新角色的权限
func (casbinService CasbinService) UpdateCasbin(authorityId string, casbinInfos []dto.PermissionInfo) error {
	casbinService.ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	success, _ := global.CASBIN.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

// 修改 casbin 的api
func (casbinService CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// 获取角色的权限列表
func (casbinService CasbinService) GetPolicyPathByRoleId(authorityId string) (pathMaps []dto.PermissionInfo) {
	list := global.CASBIN.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, dto.PermissionInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// 清除匹配的权限
func (casbinService CasbinService) ClearCasbin(v int, p ...string) bool {
	success, _ := global.CASBIN.RemoveFilteredPolicy(v, p...)
	return success
}
