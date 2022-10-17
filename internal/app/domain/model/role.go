package model

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
	Menus     []*Menu        `gorm:"many2many:role_menus;" json:"menus"` // 角色菜单多对多关系
}
