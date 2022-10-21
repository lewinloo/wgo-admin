package service

import (
	"gin_template/internal/app/domain/model"
	"gin_template/internal/app/domain/repository"
	"gin_template/internal/app/domain/service/dto"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

var MenuSet = wire.NewSet(wire.Struct(new(MenuService), "*"))

type MenuService struct {
	MenuRepo *repository.MenuRepository
	RoleRepo *repository.RoleRepository
	Log      *zap.Logger
}

// 创建菜单
func (s MenuService) Create(params dto.CreateMenuParams) bool {
	var menu model.Menu
	err := copier.Copy(&menu, params)
	if err != nil {
		s.Log.Error("copier error:" + err.Error())
		return false
	}

	_, err = s.MenuRepo.Create(menu)
	if err != nil {
		s.Log.Error("create menu error:" + err.Error())
		return false
	}

	return true
}

// 菜单列表
func (s MenuService) GetMenuTree() ([]*model.Menu, error) {
	all, err := s.MenuRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return s.formatList(all), nil
}

func (s MenuService) formatList(list []*model.Menu) []*model.Menu {
	var res []*model.Menu
	for _, menu := range list {
		// 如果父id为0则直接添加到res
		if menu.ParentId == 0 {
			res = append(res, menu)
			continue
		}
		// 父id不为0，则在res查找他的父亲节点，找到了则在他父亲节点children添加数据
		parentMenu, isExist := s.findMenuInList(res, menu.ParentId)
		if isExist {
			parentMenu.Children = append(parentMenu.Children, menu)
		}
	}

	return res
}

// 在结果集中查询
func (s MenuService) findMenuInList(list []*model.Menu, parentId uint) (*model.Menu, bool) {
	for _, menu := range list {
		if menu.ID == parentId {
			return menu, true
		}
	}
	return nil, false
}

// 删除菜单
func (s MenuService) DeleteMenu(id int) error {
	return s.MenuRepo.DeleteById(id)
}

// 查找菜单树
func (s MenuService) findMenuTree(id uint) *model.Menu {
	stack, err := s.GetMenuTree()
	if err != nil {
		return nil
	}

	for len(stack) != 0 {
		stack = stack[:len(stack)-1]
		item := stack[len(stack)-1]
		if item.ID == id {
			return item
		}
		if len(item.Children) != 0 {
			for i := 0; i < len(item.Children); i++ {
				cItem := item.Children[i]
				stack = append(stack, cItem)
			}
		}
	}

	return nil
}
