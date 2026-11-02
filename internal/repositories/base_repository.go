package repositories

import (
	"gorm.io/gorm"
)

type IBaseRepository[T any] interface {
	Create(entity *T, tx ...*gorm.DB) error
	FindByID(id interface{}, tx ...*gorm.DB) (*T, error)
	FindAll(tx ...*gorm.DB) ([]T, error)
	Update(entity *T, tx ...*gorm.DB) error
	Delete(id interface{}, tx ...*gorm.DB) error
	Count(tx ...*gorm.DB) (int64, error)
	Transaction(fn func(tx *gorm.DB) error) error
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) getDB(tx ...*gorm.DB) *gorm.DB {
	if len(tx) > 0 && tx[0] != nil {
		return tx[0]
	}
	return r.db
}

func (r *BaseRepository[T]) GetDB(tx ...*gorm.DB) *gorm.DB {
	return r.getDB(tx...)
}

func (r *BaseRepository[T]) Transaction(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *BaseRepository[T]) Create(entity *T, tx ...*gorm.DB) error {
	return r.getDB(tx...).Create(entity).Error
}

func (r *BaseRepository[T]) FindByID(id interface{}, tx ...*gorm.DB) (*T, error) {
	var entity T
	err := r.getDB(tx...).First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindAll(tx ...*gorm.DB) ([]T, error) {
	entities := make([]T, 0)
	err := r.getDB(tx...).Find(&entities).Error
	return entities, err
}

func (r *BaseRepository[T]) Update(entity *T, tx ...*gorm.DB) error {
	return r.getDB(tx...).Save(entity).Error
}

func (r *BaseRepository[T]) Delete(id interface{}, tx ...*gorm.DB) error {
	var entity T
	return r.getDB(tx...).Delete(&entity, "id = ?", id).Error
}

func (r *BaseRepository[T]) Count(tx ...*gorm.DB) (int64, error) {
	var count int64
	var entity T
	err := r.getDB(tx...).Model(&entity).Count(&count).Error
	return count, err
}
