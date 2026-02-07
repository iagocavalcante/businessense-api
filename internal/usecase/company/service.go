package company

import (
	"time"

	"businessense/internal/domain"
	"businessense/internal/repository"
)

type Service struct {
	repo repository.CompanyRepository
}

func NewService(repo repository.CompanyRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(name string, industryID, userID uint) (*domain.Company, error) {
	company := &domain.Company{
		Name:       name,
		IndustryID: industryID,
		UserID:     userID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.repo.Create(company); err != nil {
		return nil, err
	}

	return company, nil
}

func (s *Service) GetByID(id uint) (*domain.Company, error) {
	return s.repo.FindByID(id)
}

func (s *Service) ListByUser(userID uint) ([]domain.Company, error) {
	return s.repo.FindByUserID(userID)
}
