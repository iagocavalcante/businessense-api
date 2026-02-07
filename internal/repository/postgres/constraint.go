package postgres

import (
	"businessense/internal/domain"
	"businessense/internal/repository"

	"gorm.io/gorm"
)

type constraintRepository struct {
	db *gorm.DB
}

func NewConstraintRepository(db *gorm.DB) repository.ConstraintRepository {
	return &constraintRepository{db: db}
}

// FindByIssueAndSolution finds constraints associated with a specific issue-solution pair
// through the constraint_issues_solution_maps table.
func (r *constraintRepository) FindByIssueAndSolution(issueID, solutionID uint) ([]domain.Constraint, error) {
	var models []Constraint

	err := r.db.Table("constraints").
		Select("DISTINCT constraints.*").
		Joins("INNER JOIN constraint_issues_solution_maps ON constraints.id = constraint_issues_solution_maps.constraint_id").
		Where("constraint_issues_solution_maps.issue_id = ? AND constraint_issues_solution_maps.solution_id = ?", issueID, solutionID).
		Scan(&models).Error

	if err != nil {
		return nil, err
	}

	return toDomainConstraints(models), nil
}

func toDomainConstraint(m *Constraint) domain.Constraint {
	return domain.Constraint{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func toDomainConstraints(models []Constraint) []domain.Constraint {
	result := make([]domain.Constraint, len(models))
	for i, m := range models {
		result[i] = toDomainConstraint(&m)
	}
	return result
}
