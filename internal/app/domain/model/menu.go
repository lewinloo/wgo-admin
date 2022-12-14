package model

import (
	"gin_template/internal/app/common"
)

type Menu struct {
	common.BaseModel
	Name              string  `gorm:"type:varchar(50);comment:'菜单名称(英文名, 可用于国际化)'" json:"name"`
	Title             string  `gorm:"type:varchar(50);comment:'菜单标题(无法国际化时使用)'" json:"title"`
	Icon              string  `gorm:"type:varchar(50);comment:'菜单图标'" json:"icon"`
	Path              string  `gorm:"type:varchar(100);comment:'菜单访问路径'" json:"path"`
	Redirect          string  `gorm:"type:varchar(100);comment:'重定向路径'" json:"redirect"`
	ComponentFilePath string  `gorm:"type:varchar(100);comment:'前端组件路径'" json:"componentFilePath"`
	Sort              uint    `gorm:"type:int(3) unsigned;default:999;comment:'菜单顺序(1-999)'" json:"sort"`
	Enable            bool    `gorm:"default:1;comment:'菜单状态(正常/禁用, 默认正常)'" json:"enable"`
	Hidden            bool    `gorm:"default:0;comment:'菜单在侧边栏隐藏'" json:"hidden"`
	KeepAlive         bool    `gorm:"default:1;comment:'菜单是否被 <keep-alive> '" json:"keepAlive"`
	Breadcrumb        bool    `gorm:"default:1;comment:'面包屑可见性(可见/隐藏, 默认可见)'" json:"breadcrumb"`
	ParentId          uint    `gorm:"default:0;comment:'父菜单编号(编号为0时表示根菜单)'" json:"parentId"`
	Children          []*Menu `gorm:"-" json:"children"`                  // 子菜单集合
	Roles             []*Role `gorm:"many2many:role_menus;" json:"roles"` // 角色菜单多对多关系
}
