package models

import (
	"github.com/jinzhu/gorm"
)

//ConstraintIssuesSolutionMap Type
type ConstraintIssuesSolutionMap struct {
	gorm.Model
	Constraint   Constraint
	ConstraintID int
	Issue        Issue
	IssueID      int
	Solution     Solution
	SolutionID   int
	Weight       float64
}
