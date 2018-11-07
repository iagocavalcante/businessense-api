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

//Create IssuePainPointsMap
func (issuepainpointmap *IssuePainPointsMap) Create() map[string]interface{} {
	GetDB().Create(issuepainpointmap)
	response := u.Message(true, "Issue has been created")
	response["issuepainpointmap"] = issuepainpointmap
	return response
}
