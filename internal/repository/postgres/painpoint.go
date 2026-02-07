package postgres

import (
	"businessense/internal/domain"
	"businessense/internal/repository"

	"gorm.io/gorm"
)

type painPointRepository struct {
	db *gorm.DB
}

func NewPainPointRepository(db *gorm.DB) repository.PainPointRepository {
	return &painPointRepository{db: db}
}

func (r *painPointRepository) Create(pp *domain.PainPoint) error {
	model := &PainPoint{
		Name: pp.Name,
	}
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	pp.ID = model.ID
	pp.CreatedAt = model.CreatedAt
	pp.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *painPointRepository) FindAll() ([]domain.PainPoint, error) {
	var models []PainPoint
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	return toDomainPainPoints(models), nil
}

func (r *painPointRepository) Search(query string) ([]domain.PainPoint, error) {
	var models []PainPoint
	if err := r.db.Where("name ILIKE ?", "%"+query+"%").Find(&models).Error; err != nil {
		return nil, err
	}
	return toDomainPainPoints(models), nil
}

func toDomainPainPoint(m *PainPoint) domain.PainPoint {
	return domain.PainPoint{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func toDomainPainPoints(models []PainPoint) []domain.PainPoint {
	result := make([]domain.PainPoint, len(models))
	for i, m := range models {
		result[i] = toDomainPainPoint(&m)
	}
	return result
}
