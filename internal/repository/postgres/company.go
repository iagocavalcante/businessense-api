package postgres

import (
	"businessense/internal/domain"
	"businessense/internal/repository"

	"gorm.io/gorm"
)

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) repository.CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) Create(company *domain.Company) error {
	model := &Company{
		Name:       company.Name,
		IndustryID: company.IndustryID,
		UserID:     company.UserID,
	}
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	company.ID = model.ID
	company.CreatedAt = model.CreatedAt
	company.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *companyRepository) FindByID(id uint) (*domain.Company, error) {
	var model Company
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return toDomainCompany(&model), nil
}

func (r *companyRepository) FindByUserID(userID uint) ([]domain.Company, error) {
	var models []Company
	if err := r.db.Where("user_id = ?", userID).Find(&models).Error; err != nil {
		return nil, err
	}
	return toDomainCompanies(models), nil
}

func toDomainCompany(m *Company) *domain.Company {
	return &domain.Company{
		ID:         m.ID,
		Name:       m.Name,
		IndustryID: m.IndustryID,
		UserID:     m.UserID,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func toDomainCompanies(models []Company) []domain.Company {
	result := make([]domain.Company, len(models))
	for i, m := range models {
		result[i] = *toDomainCompany(&m)
	}
	return result
}
