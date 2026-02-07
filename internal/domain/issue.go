package domain

import "time"

type Issue struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stats       string    `json:"stats"`
	RiskFactors string    `json:"riskFactors"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type IssueRelevance struct {
	IssueID   uint    `json:"issueId"`
	Name      string  `json:"name"`
	Relevance float64 `json:"relevance"`
}

type IssuePainPointMap struct {
	ID          uint `json:"id"`
	IssueID     uint `json:"issueId"`
	PainPointID uint `json:"painPointId"`
}
