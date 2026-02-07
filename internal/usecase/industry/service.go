package industry

import (
	"time"

	"businessense/internal/domain"
	"businessense/internal/repository"
)

type Service struct {
	repo repository.IndustryRepository
}

func NewService(repo repository.IndustryRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(name string) (*domain.Industry, error) {
	industry := &domain.Industry{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(industry); err != nil {
		return nil, err
	}

	return industry, nil
}

func (s *Service) List() ([]domain.Industry, error) {
	return s.repo.FindAll()
}

func (s *Service) Search(query string) ([]domain.Industry, error) {
	return s.repo.Search(query)
}
