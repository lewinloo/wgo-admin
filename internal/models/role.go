package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        string         `gorm:"not null;unique;primaryKey;index;comment:角色id;size:90" json:"id"` // 角色id
	CreatedAt time.Time      `json:"createdAt"`                                                       // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`                                                       // 创建时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                                  // 删除时间
	Name      string         `gorm:"type:varchar(20);not null;unique" json:"name"`
	Desc      string         `gorm:"type:varchar(100);" json:"desc"`
	Status    uint           `gorm:"type:tinyint(1);default:1;comment:'1正常, 2禁用'" json:"status"`
	Users     []*User        `gorm:"many2many:user_roles" json:"users"`

	// Sort   uint    `gorm:"type:int(3);default:999;comment:'角色排序(排序越大权限越低, 不能查看比自己序号小的角色, 不能编辑同序号用户权限, 排序为1表示超级管理员)'" json:"sort"`
	// Menus   []*Menu `gorm:"many2many:role_menus;" json:"menus"` // 角色菜单多对多关系
}
