package service

import (
	"gin_template/internal/app/domain/model"
	"gin_template/internal/app/domain/repository"
	"gin_template/internal/app/domain/service/dto"
	"gin_template/internal/app/domain/service/vo"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleService), "*"))

type RoleService struct {
	MenuRepo *repository.MenuRepository
	RoleRepo *repository.RoleRepository
	Log      *zap.Logger
}

// 创建角色
func (s RoleService) Create(params dto.CreateRoleParams) bool {
	var role model.Role
	err := copier.Copy(&role, params)
	if err != nil {
		s.Log.Error("copier error:" + err.Error())
		return false
	}

	_, err = s.RoleRepo.Create(role)
	if err != nil {
		s.Log.Error("create role error:" + err.Error())
		return false
	}

	return true
}

// 删除角色
func (s RoleService) DeleteById(id int) error {
	return s.RoleRepo.DeleteById(id)
}

// 查询角色列表
func (s RoleService) FindList(params dto.QueryRoleListParams) (vo.ListResult, error) {
	return s.RoleRepo.FindList(params)
}
