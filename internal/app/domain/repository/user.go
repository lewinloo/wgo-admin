package repository

import (
	"gin_template/internal/app/domain/model"
	"gin_template/internal/app/domain/service/dto"
	"gin_template/internal/app/domain/service/vo"
	"github.com/google/wire"
	"gorm.io/gorm"
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
	FindList(params dto.QueryUserListParams) (vo.ListResult, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (u UserRepository) FindList(params dto.QueryUserListParams) (vo.ListResult, error) {
	if params.Current == 0 {
		params.Current = 1
	}

	if params.Size == 0 {
		params.Size = 10
	}
	limit := params.Current
	offset := params.Size * (params.Current - 1)
	db := u.DB.Model(&model.User{})
	var total int64
	var userList []model.User

	if params.Username != "" {
		db = db.Where("username LIKE ?", "%"+params.Username+"%")
	}

	if params.Email != "" {
		db = db.Where("email LIKE ?", "%"+params.Email+"%")
	}

	err := db.Count(&total).Error

	if err != nil {
		return vo.ListResult{}, err
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Find(&userList).Error
	}

	res := vo.ListResult{
		List:    userList,
		Total:   total,
		Current: params.Current,
		Size:    params.Size,
	}

	return res, nil
}

func (u UserRepository) Create(model model.User) (bool, error) {
	err := u.DB.Create(&model).Error
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (u UserRepository) FindById(id int) (model model.User, err error) {
	return model, u.DB.First(&model, id).Error
}

func (u UserRepository) FindAll() (list []model.User, err error) {
	return list, u.DB.Find(&list).Error
}

func (u UserRepository) Find(condition ...any) (list []model.User, err error) {
	return list, u.DB.Find(&list, condition).Error
}

func (u UserRepository) FindOne(condition ...any) (model model.User, err error) {
	return model, u.DB.First(&model).Where(condition).Error
}

func (u UserRepository) DeleteById(id int) (bool, error) {
	var model model.User
	if err := u.DB.Delete(&model, "WHERE id = ?", id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (u UserRepository) DeleteByIds(ids []int) (bool, error) {
	var models []model.User
	if err := u.DB.Delete(&models, "id in ?", ids).Error; err != nil {
		return false, err
	}
	return true, nil
}
