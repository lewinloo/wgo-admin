package repository

import (
	"gin_template/internal/app/domain/model"
	"gin_template/internal/app/domain/service/dto"
	"gin_template/internal/app/domain/service/vo"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RoleSet = wire.NewSet(
	wire.Struct(new(RoleRepository), "*"),
	wire.Bind(new(IRoleRepository), new(*RoleRepository)),
)

type IRoleRepository interface {
	Create(model model.Role) (bool, error)
	FindById(id int) (model model.Role, err error)
	FindAll() (list []model.Role, err error)
	Find(condition ...any) (list []model.Role, err error)
	FindOne(condition ...any) (model model.Role, err error)
	DeleteById(id int) error
	DeleteByIds(ids []int) (bool, error)
	FindList(params dto.QueryRoleListParams) (result vo.ListResult, err error)
}

type RoleRepository struct {
	DB *gorm.DB
}

func (r RoleRepository) FindList(params dto.QueryRoleListParams) (result vo.ListResult, err error) {
	if params.Current == 0 {
		params.Current = 1
	}

	if params.Size == 0 {
		params.Size = 10
	}
	limit := params.Current
	offset := params.Size * (params.Current - 1)
	db := r.DB.Model(&model.Role{})
	var total int64
	var roleList []model.Role

	err = db.Count(&total).Error

	if err != nil {
		return vo.ListResult{}, err
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Find(&roleList).Error
	}

	res := vo.ListResult{
		List:    roleList,
		Total:   total,
		Current: params.Current,
		Size:    params.Size,
	}
	return res, nil

}

func (r RoleRepository) Create(model model.Role) (bool, error) {
	err := r.DB.Create(&model).Error
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (r RoleRepository) FindById(id int) (model model.Role, err error) {
	return model, r.DB.First(&model, id).Error
}

func (r RoleRepository) FindAll() (list []model.Role, err error) {
	return list, r.DB.Find(&list).Error
}

func (r RoleRepository) Find(condition ...any) (list []model.Role, err error) {
	return list, r.DB.Find(&list, condition).Error
}

func (r RoleRepository) FindOne(condition ...any) (model model.Role, err error) {
	return model, r.DB.First(&model, condition).Error
}

func (r RoleRepository) DeleteById(id int) error {
	var role model.Role
	var err error
	db := r.DB.Preload("Users").Preload("Menus").Where("id = ?", id).First(&role).Delete(&role)

	err = db.Error

	// 解除绑定用户的关系
	if len(role.Users) > 0 {
		err = r.DB.Model(&role).Association("Users").Delete(&role.Users)
	} else {
		err = db.Error
	}

	// 解除绑定菜单的关系
	if len(role.Menus) > 0 {
		err = r.DB.Model(&role).Association("Menus").Delete(&role.Menus)
	} else {
		err = db.Error
	}

	return err
}

func (r RoleRepository) DeleteByIds(ids []int) (bool, error) {
	var models []model.Role
	if err := r.DB.Delete(&models, "id in ?", ids).Error; err != nil {
		return false, err
	}
	return true, nil
}
