package postgres

import (
	"businessense/internal/domain"
	"businessense/internal/repository"

	"gorm.io/gorm"
)

type industryRepository struct {
	db *gorm.DB
}

func NewIndustryRepository(db *gorm.DB) repository.IndustryRepository {
	return &industryRepository{db: db}
}

func (r *industryRepository) Create(industry *domain.Industry) error {
	model := &Industry{
		Name: industry.Name,
	}
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	industry.ID = model.ID
	industry.CreatedAt = model.CreatedAt
	industry.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *industryRepository) FindAll() ([]domain.Industry, error) {
	var models []Industry
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	return toDomainIndustries(models), nil
}

func (r *industryRepository) Search(query string) ([]domain.Industry, error) {
	var models []Industry
	if err := r.db.Where("name ILIKE ?", "%"+query+"%").Find(&models).Error; err != nil {
		return nil, err
	}
	return toDomainIndustries(models), nil
}

func toDomainIndustry(m *Industry) domain.Industry {
	return domain.Industry{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func toDomainIndustries(models []Industry) []domain.Industry {
	result := make([]domain.Industry, len(models))
	for i, m := range models {
		result[i] = toDomainIndustry(&m)
	}
	return result
}
