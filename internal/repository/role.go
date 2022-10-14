package repository

import (
	"gin_template/internal/models"
	"sync"
)

var (
	Role     *RoleRepository
	onceRole sync.Once
)

// 单例模式
func NewRole() *RoleRepository {
	onceRole.Do(func() {
		Role = &RoleRepository{}
	})
	return Role
}

// 扩展基础的Repository
type RoleRepository struct {
	Repository[models.Role]
}
