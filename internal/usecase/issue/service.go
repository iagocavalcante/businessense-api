package issue

import (
	"businessense/internal/domain"
	"businessense/internal/repository"
)

type Service struct {
	repo repository.IssueRepository
}

func NewService(repo repository.IssueRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetByID(id uint) (*domain.Issue, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetRelevance(painPointIDs []uint) ([]domain.IssueRelevance, error) {
	return s.repo.FindByPainPointIDs(painPointIDs)
}
