package repository

import (
	"gin_template/internal/models"
)

var (
	Role *RoleRepository
)

// 单例模式
func NewRole() *RoleRepository {
	once.Do(func() {
		Role = &RoleRepository{}
	})
	return Role
}

// 扩展基础的Repository
type RoleRepository struct {
	Repository[models.Role]
}
