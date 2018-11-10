package models

import (
	u "businessense/utils"

	"github.com/jinzhu/gorm"
)

//IssuePainPointsMap Type
type IssuePainPointsMap struct {
	gorm.Model
	Issue       Issue
	IssueID     int
	PainPoint   PainPoint
	PainPointID int
	Relevance   float64
}

type IssueRelevance struct {
	IssueID   int
	Name      string
	Relevance float64
}

//Create IssuePainPointsMap
func (issuepainpointmap *IssuePainPointsMap) Create() map[string]interface{} {
	GetDB().Create(issuepainpointmap)
	response := u.Message(true, "Issue has been created")
	response["issuepainpointmap"] = issuepainpointmap
	return response
}

//
func GetIssuesRelevance(painpoints []string) []*IssueRelevance {
	s := "select issue_id, issues.name, sum(relevance) as relevance " +
		"from issue_pain_points_maps " +
		"inner join issues on issues.id = issue_id " +
		"where pain_point_id in (?) " +
		"group by issue_id, issues.name "
	issuesRelevance := make([]*IssueRelevance, 0)
	GetDB().Raw(s, painpoints).Scan(&issuesRelevance)
	return issuesRelevance
}
