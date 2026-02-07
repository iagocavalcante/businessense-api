package postgres

import (
	"businessense/internal/domain"
	"businessense/internal/repository"

	"gorm.io/gorm"
)

type solutionRepository struct {
	db *gorm.DB
}

func NewSolutionRepository(db *gorm.DB) repository.SolutionRepository {
	return &solutionRepository{db: db}
}

func (r *solutionRepository) FindByID(id uint) (*domain.Solution, error) {
	var model Solution
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return toDomainSolution(&model), nil
}

// FindByIssueID finds solutions associated with a given issue
// through the constraint_issues_solution_maps table.
func (r *solutionRepository) FindByIssueID(issueID uint) ([]domain.Solution, error) {
	var models []Solution

	err := r.db.Table("solutions").
		Select("DISTINCT solutions.*").
		Joins("INNER JOIN constraint_issues_solution_maps ON solutions.id = constraint_issues_solution_maps.solution_id").
		Where("constraint_issues_solution_maps.issue_id = ?", issueID).
		Scan(&models).Error

	if err != nil {
		return nil, err
	}

	return toDomainSolutions(models), nil
}

func toDomainSolution(m *Solution) *domain.Solution {
	return &domain.Solution{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func toDomainSolutions(models []Solution) []domain.Solution {
	result := make([]domain.Solution, len(models))
	for i, m := range models {
		result[i] = *toDomainSolution(&m)
	}
	return result
}
