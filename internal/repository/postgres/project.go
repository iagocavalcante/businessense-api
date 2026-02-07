package postgres

import (
	"businessense/internal/domain"
	"businessense/internal/repository"

	"gorm.io/gorm"
)

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) repository.ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Create(project *domain.Project) error {
	model := &Project{
		Name:        project.Name,
		Description: project.Description,
		CompanyID:   project.CompanyID,
	}
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	project.ID = model.ID
	project.CreatedAt = model.CreatedAt
	project.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *projectRepository) FindByID(id uint) (*domain.Project, error) {
	var model Project
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return toDomainProject(&model), nil
}

func (r *projectRepository) FindByCompanyID(companyID uint) ([]domain.Project, error) {
	var models []Project
	if err := r.db.Where("company_id = ?", companyID).Find(&models).Error; err != nil {
		return nil, err
	}
	return toDomainProjects(models), nil
}

func toDomainProject(m *Project) *domain.Project {
	return &domain.Project{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		CompanyID:   m.CompanyID,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func toDomainProjects(models []Project) []domain.Project {
	result := make([]domain.Project, len(models))
	for i, m := range models {
		result[i] = *toDomainProject(&m)
	}
	return result
}
