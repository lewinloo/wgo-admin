package repository

import (
	"errors"
	"fmt"
	"gin_template/internal/app/domain/model"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var MenuSet = wire.NewSet(
	wire.Struct(new(MenuRepository), "*"),
	wire.Bind(new(IMenuRepository), new(*MenuRepository)),
)

type IMenuRepository interface {
	Create(model model.Menu) (bool, error)
	FindById(id int) (model *model.Menu, err error)
	FindAll() (list []*model.Menu, err error)
	Find(condition ...any) (list []*model.Menu, err error)
	FindOne(condition ...any) (model *model.Menu, err error)
	DeleteById(id int) error
	DeleteByIds(ids []int) (bool, error)
}

type MenuRepository struct {
	DB *gorm.DB
}

func (r MenuRepository) Create(model model.Menu) (bool, error) {
	err := r.DB.Create(&model).Error
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (r MenuRepository) FindById(id int) (model *model.Menu, err error) {
	return model, r.DB.First(&model, id).Error
}

func (r MenuRepository) FindAll() (list []*model.Menu, err error) {
	return list, r.DB.Find(&list).Error
}

func (r MenuRepository) Find(condition ...any) (list []*model.Menu, err error) {
	return list, r.DB.Find(&list, condition).Error
}

func (r MenuRepository) FindOne(condition ...any) (model *model.Menu, err error) {
	return model, r.DB.First(&model, condition).Error
}

func (r MenuRepository) DeleteById(id int) error {
	err := r.DB.Where("parent_id = ?", id).First(&model.Menu{}).Error
	if err != nil {
		// 查询menu并预加载绑定的角色信息
		var menu model.Menu
		db := r.DB.Preload("Roles").Where("id = ?", id).First(&menu).Delete(&menu)
		fmt.Println(menu)
		// 如果改菜单绑定了角色，则解除与改角色的关系
		if len(menu.Roles) > 0 {
			err = r.DB.Model(&menu).Association("Roles").Delete(&menu.Roles)
		} else {
			err = db.Error
		}
		return err
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
}

func (r MenuRepository) DeleteByIds(ids []int) (bool, error) {
	var models []model.Menu
	if err := r.DB.Delete(&models, "id in ?", ids).Error; err != nil {
		return false, err
	}
	return true, nil
}
