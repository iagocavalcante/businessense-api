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

//GetSolutionsForIssue based on issue id
func GetConstraintsForIssue(issueid int) []*ConstraintIssuesSolutionMap {
	constraintIssueSolutions := make([]*ConstraintIssuesSolutionMap, 0)
	GetDB().Set("gorm:auto_preload", true).Find(&constraintIssueSolutions).Where("issue_id = ?)", issueid)
	return constraintIssueSolutions
}
