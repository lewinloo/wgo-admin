package repository

import "gin_template/internal/global"

func New[T any](model T) *Repository[T] {
	return new(Repository[T])
}

type BaseReposity[T any] interface {
	Create(model *T) (bool, error)
	FindById(id int) (T, error)
	FindAll() (list []T, err error)
	DeleteById(id int) (bool, error)
	DeleteByIds(ids []int) (bool, error)
}

type Repository[T any] struct{}

func (r *Repository[T]) Create(model *T) (bool, error) {
	err := global.DB.Create(model).Error
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (r *Repository[T]) FindById(id int) (T, error) {
	var model T
	if err := global.DB.First(&model, "WHERE id = ?", id).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (r *Repository[T]) FindAll() (list []T, err error) {
	return list, global.DB.Find(&list).Error
}

func (r *Repository[T]) DeleteById(id int) (bool, error) {
	var model T
	if err := global.DB.Delete(&model, "WHERE id = ?", id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *Repository[T]) DeleteByIds(ids int) (bool, error) {
	var models []T
	if err := global.DB.Delete(&models, "id in ?", ids).Error; err != nil {
		return false, err
	}
	return true, nil
}
