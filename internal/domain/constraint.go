package domain

import "time"

type Constraint struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ConstraintIssueSolutionMap struct {
	ID           uint `json:"id"`
	ConstraintID uint `json:"constraintId"`
	IssueID      uint `json:"issueId"`
	SolutionID   uint `json:"solutionId"`
}
