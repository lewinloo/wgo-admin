package model

import "gin_template/internal/app/common"

type Resource struct {
	common.BaseModel
	Name   string `json:"name" gorm:"comment:资源名称"`
	Perms  string `json:"perms" gorm:"comment:权限标识"`
	Method string `json:"method" gorm:"default:POST;comment:资源方法"`
	Path   string `json:"path" gorm:"comment:资源路径"`
	Group  string `json:"group" gorm:"comment:资源组"`
}
