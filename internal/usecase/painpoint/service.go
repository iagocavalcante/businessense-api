package painpoint

import (
	"time"

	"businessense/internal/domain"
	"businessense/internal/repository"
)

type Service struct {
	repo repository.PainPointRepository
}

func NewService(repo repository.PainPointRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(name string) (*domain.PainPoint, error) {
	painPoint := &domain.PainPoint{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(painPoint); err != nil {
		return nil, err
	}

	return painPoint, nil
}

func (s *Service) List() ([]domain.PainPoint, error) {
	return s.repo.FindAll()
}

func (s *Service) Search(query string) ([]domain.PainPoint, error) {
	return s.repo.Search(query)
}
