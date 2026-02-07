package solution

import (
	"businessense/internal/domain"
	"businessense/internal/repository"
)

type Service struct {
	solutionRepo   repository.SolutionRepository
	constraintRepo repository.ConstraintRepository
}

func NewService(solutionRepo repository.SolutionRepository, constraintRepo repository.ConstraintRepository) *Service {
	return &Service{
		solutionRepo:   solutionRepo,
		constraintRepo: constraintRepo,
	}
}

func (s *Service) GetByID(id uint) (*domain.Solution, error) {
	return s.solutionRepo.FindByID(id)
}

func (s *Service) GetByIssue(issueID uint) ([]domain.Solution, error) {
	return s.solutionRepo.FindByIssueID(issueID)
}

func (s *Service) GetConstraints(issueID, solutionID uint) ([]domain.Constraint, error) {
	return s.constraintRepo.FindByIssueAndSolution(issueID, solutionID)
}
