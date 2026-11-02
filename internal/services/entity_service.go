package services

import (
	"errors"

	"api-puri/internal/dto"
	"api-puri/internal/models"
	"api-puri/internal/repositories"
	"gorm.io/gorm"
)

var (
	ErrEntityNotFound   = errors.New("entity not found")
	ErrEntityCodeExists = errors.New("entity code already exists")
)

type EntityService interface {
	Create(req *dto.CreateEntityRequest) (*dto.EntityResponse, error)
	Update(id uint, req *dto.UpdateEntityRequest) (*dto.EntityResponse, error)
	Delete(id uint) error
	GetByID(id uint) (*dto.EntityResponse, error)
	GetPaginated(page int, perPage int, search string, category string, status string) ([]dto.EntityResponse, int64, error)
	GetMapPoints(category string, status string) ([]dto.EntityMapPointDTO, error)
}

type entityService struct {
	repo repositories.EntityRepository
}

func NewEntityService(repo repositories.EntityRepository) EntityService {
	return &entityService{repo: repo}
}

func (s *entityService) Create(req *dto.CreateEntityRequest) (*dto.EntityResponse, error) {
	existing, err := s.repo.FindByCode(req.Code)
	if err == nil && existing != nil {
		return nil, ErrEntityCodeExists
	}

	metadata := models.JSONMetadata(req.Metadata)
	if metadata == nil {
		metadata = make(models.JSONMetadata)
	}

	entity := &models.Entity{
		Name:        req.Name,
		Code:        req.Code,
		Category:    models.EntityCategory(req.Category),
		Status:      models.EntityStatus(req.Status),
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Description: req.Description,
		Metadata:    metadata,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	return s.toResponse(entity), nil
}

func (s *entityService) Update(id uint, req *dto.UpdateEntityRequest) (*dto.EntityResponse, error) {
	entity, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrEntityNotFound
		}
		return nil, err
	}

	entity.Name = req.Name
	entity.Category = models.EntityCategory(req.Category)
	entity.Status = models.EntityStatus(req.Status)
	entity.Latitude = req.Latitude
	entity.Longitude = req.Longitude
	entity.Description = req.Description

	if req.Metadata != nil {
		entity.Metadata = models.JSONMetadata(req.Metadata)
	}

	if err := s.repo.Update(entity); err != nil {
		return nil, err
	}

	return s.toResponse(entity), nil
}

func (s *entityService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrEntityNotFound
		}
		return err
	}
	return s.repo.Delete(id)
}

func (s *entityService) GetByID(id uint) (*dto.EntityResponse, error) {
	entity, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrEntityNotFound
		}
		return nil, err
	}
	return s.toResponse(entity), nil
}

func (s *entityService) GetPaginated(page int, perPage int, search string, category string, status string) ([]dto.EntityResponse, int64, error) {
	entities, total, err := s.repo.FindPaginated(page, perPage, search, category, status)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]dto.EntityResponse, 0, len(entities))
	for i := range entities {
		responses = append(responses, *s.toResponse(&entities[i]))
	}

	return responses, total, nil
}

func (s *entityService) GetMapPoints(category string, status string) ([]dto.EntityMapPointDTO, error) {
	entities, err := s.repo.FindMapPoints(category, status)
	if err != nil {
		return nil, err
	}

	points := make([]dto.EntityMapPointDTO, 0, len(entities))
	for _, e := range entities {
		points = append(points, dto.EntityMapPointDTO{
			ID:        e.ID,
			Name:      e.Name,
			Code:      e.Code,
			Category:  string(e.Category),
			Status:    string(e.Status),
			Latitude:  e.Latitude,
			Longitude: e.Longitude,
		})
	}

	return points, nil
}

func (s *entityService) toResponse(e *models.Entity) *dto.EntityResponse {
	return &dto.EntityResponse{
		ID:          e.ID,
		Name:        e.Name,
		Code:        e.Code,
		Category:    string(e.Category),
		Status:      string(e.Status),
		Latitude:    e.Latitude,
		Longitude:   e.Longitude,
		Description: e.Description,
		Metadata:    map[string]interface{}(e.Metadata),
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
