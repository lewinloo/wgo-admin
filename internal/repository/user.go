package repository

import (
	"gin_template/internal/models"
	"sync"
)

var (
	User     *UserRepository
	onceUser sync.Once
)

// 单例模式
func NewUser() *UserRepository {
	onceUser.Do(func() {
		User = &UserRepository{}
	})
	return User
}

// 扩展基础的Repository
type UserRepository struct {
	Repository[models.User]
}
