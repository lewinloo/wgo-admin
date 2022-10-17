package repository

import (
	"gin_template/internal/app/domain/model"
	"gin_template/internal/app/global"
)

func NewRole() RoleRepository {
	return RoleRepository{}
}

type IRoleRepository interface {
	Create(model model.Role) (bool, error)
	FindById(id int) (model model.Role, err error)
	FindAll() (list []model.Role, err error)
	Find(condition ...any) (list []model.Role, err error)
	FindOne(condition ...any) (model model.Role, err error)
	DeleteById(id int) (bool, error)
	DeleteByIds(ids []int) (bool, error)
}

type RoleRepository struct {
}

func (r RoleRepository) Create(model model.Role) (bool, error) {
	err := global.DB.Create(&model).Error
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (r RoleRepository) FindById(id int) (model model.Role, err error) {
	return model, global.DB.First(&model, id).Error
}

func (r RoleRepository) FindAll() (list []model.Role, err error) {
	return list, global.DB.Find(&list).Error
}

func (r RoleRepository) Find(condition ...any) (list []model.Role, err error) {
	return list, global.DB.Find(&list, condition).Error
}

func (r RoleRepository) FindOne(condition ...any) (model model.Role, err error) {
	return model, global.DB.First(&model, condition).Error
}

func (r RoleRepository) DeleteById(id int) (bool, error) {
	var model model.Role
	if err := global.DB.Delete(&model, "WHERE id = ?", id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r RoleRepository) DeleteByIds(ids []int) (bool, error) {
	var models []model.Role
	if err := global.DB.Delete(&models, "id in ?", ids).Error; err != nil {
		return false, err
	}
	return true, nil
}
