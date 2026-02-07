package project

import (
	"time"

	"businessense/internal/domain"
	"businessense/internal/repository"
)

type Service struct {
	repo repository.ProjectRepository
}

func NewService(repo repository.ProjectRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(name, description string, companyID uint) (*domain.Project, error) {
	project := &domain.Project{
		Name:        name,
		Description: description,
		CompanyID:   companyID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *Service) GetByID(id uint) (*domain.Project, error) {
	return s.repo.FindByID(id)
}

func (s *Service) ListByCompany(companyID uint) ([]domain.Project, error) {
	return s.repo.FindByCompanyID(companyID)
}
