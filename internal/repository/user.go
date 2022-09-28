package repository

import (
	"gin_template/internal/models"
	"sync"
)

var (
	User *UserRepository
	once sync.Once
)

// 单例模式
func NewUser() *UserRepository {
	once.Do(func() {
		User = &UserRepository{}
	})
	return User
}

// 扩展基础的Repository
type UserRepository struct {
	Repository[models.User]
}

func (r *UserRepository) AddTest() {
}
