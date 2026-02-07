package postgres

import (
	"businessense/internal/domain"
	"businessense/internal/repository"

	"gorm.io/gorm"
)

type issueRepository struct {
	db *gorm.DB
}

func NewIssueRepository(db *gorm.DB) repository.IssueRepository {
	return &issueRepository{db: db}
}

func (r *issueRepository) FindByID(id uint) (*domain.Issue, error) {
	var model Issue
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return toDomainIssue(&model), nil
}

// FindByPainPointIDs finds issues related to the given pain point IDs
// and calculates relevance based on how many of the provided pain points
// are associated with each issue.
func (r *issueRepository) FindByPainPointIDs(ppIDs []uint) ([]domain.IssueRelevance, error) {
	if len(ppIDs) == 0 {
		return []domain.IssueRelevance{}, nil
	}

	// Query to find issues and count how many of the provided pain points match
	type issueRelevanceResult struct {
		IssueID   uint
		Name      string
		MatchCount int
	}

	var results []issueRelevanceResult

	// Find issues linked to the provided pain points and count matches
	err := r.db.Table("issues").
		Select("issues.id as issue_id, issues.name, COUNT(issues_pain_points_maps.pain_point_id) as match_count").
		Joins("INNER JOIN issues_pain_points_maps ON issues.id = issues_pain_points_maps.issue_id").
		Where("issues_pain_points_maps.pain_point_id IN ?", ppIDs).
		Group("issues.id, issues.name").
		Order("match_count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// Calculate relevance as percentage of provided pain points that match
	totalPainPoints := float64(len(ppIDs))
	issueRelevances := make([]domain.IssueRelevance, len(results))
	for i, res := range results {
		issueRelevances[i] = domain.IssueRelevance{
			IssueID:   res.IssueID,
			Name:      res.Name,
			Relevance: float64(res.MatchCount) / totalPainPoints,
		}
	}

	return issueRelevances, nil
}

func toDomainIssue(m *Issue) *domain.Issue {
	return &domain.Issue{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Stats:       m.Stats,
		RiskFactors: m.RiskFactors,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
