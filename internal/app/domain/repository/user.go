package repository

import (
	"gin_template/internal/app/domain/model"
	"gin_template/internal/app/global"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(
	wire.Struct(new(UserRepository), "*"),
	wire.Bind(new(IUserRepository), new(*UserRepository)),
)

type IUserRepository interface {
	Create(model model.User) (bool, error)
	FindById(id int) (model model.User, err error)
	FindAll() (list []model.User, err error)
	Find(condition ...any) (list []model.User, err error)
	FindOne(condition ...any) (model model.User, err error)
	DeleteById(id int) (bool, error)
	DeleteByIds(ids []int) (bool, error)
}

type UserRepository struct{}

func (u UserRepository) Create(model model.User) (bool, error) {
	err := global.DB.Create(&model).Error
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (u UserRepository) FindById(id int) (model model.User, err error) {
	return model, global.DB.First(&model, id).Error
}

func (u UserRepository) FindAll() (list []model.User, err error) {
	return list, global.DB.Find(&list).Error
}

func (u UserRepository) Find(condition ...any) (list []model.User, err error) {
	return list, global.DB.Find(&list, condition).Error
}

func (u UserRepository) FindOne(condition ...any) (model model.User, err error) {
	return model, global.DB.First(&model).Where(condition).Error
}

func (u UserRepository) DeleteById(id int) (bool, error) {
	var model model.User
	if err := global.DB.Delete(&model, "WHERE id = ?", id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (u UserRepository) DeleteByIds(ids []int) (bool, error) {
	var models []model.User
	if err := global.DB.Delete(&models, "id in ?", ids).Error; err != nil {
		return false, err
	}
	return true, nil
}
