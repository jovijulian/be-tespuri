package repositories

import (
	"strings"

	"api-puri/internal/models"
	"gorm.io/gorm"
)

type EntityRepository interface {
	IBaseRepository[models.Entity]
	FindByCode(code string, tx ...*gorm.DB) (*models.Entity, error)
	FindPaginated(page int, perPage int, search string, category string, status string) ([]models.Entity, int64, error)
	FindMapPoints(category string, status string) ([]models.Entity, error)
}

type entityRepository struct {
	*BaseRepository[models.Entity]
}

func NewEntityRepository(db *gorm.DB) EntityRepository {
	return &entityRepository{
		BaseRepository: NewBaseRepository[models.Entity](db),
	}
}

func (r *entityRepository) FindByCode(code string, tx ...*gorm.DB) (*models.Entity, error) {
	var entity models.Entity
	err := r.getDB(tx...).Where("code = ?", code).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *entityRepository) FindPaginated(page int, perPage int, search string, category string, status string) ([]models.Entity, int64, error) {
	var entities []models.Entity
	var total int64

	query := r.db.Model(&models.Entity{})

	if search != "" {
		s := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(code) LIKE ?", s, s)
	}

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage
	err := query.Order("created_at DESC").Limit(perPage).Offset(offset).Find(&entities).Error
	if err != nil {
		return nil, 0, err
	}

	return entities, total, nil
}

func (r *entityRepository) FindMapPoints(category string, status string) ([]models.Entity, error) {
	var entities []models.Entity

	query := r.db.Model(&models.Entity{})

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Select("id", "name", "code", "category", "status", "latitude", "longitude").Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return entities, nil
}
