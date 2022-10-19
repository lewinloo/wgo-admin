package repository

import (
	"gin_template/internal/app/domain/model"
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
	DeleteById(id int) (bool, error)
	DeleteByIds(ids []int) (bool, error)
}

type RoleRepository struct {
	DB *gorm.DB
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

func (r RoleRepository) DeleteById(id int) (bool, error) {
	var model model.Role
	if err := r.DB.Delete(&model, "WHERE id = ?", id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r RoleRepository) DeleteByIds(ids []int) (bool, error) {
	var models []model.Role
	if err := r.DB.Delete(&models, "id in ?", ids).Error; err != nil {
		return false, err
	}
	return true, nil
}
