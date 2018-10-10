package models

import (
	u "businessense/utils"

	"github.com/jinzhu/gorm"
)

//IssuePainPointsMap Type
type IssuePainPointsMap struct {
	gorm.Model
	issue       Issue
	IssueID     int
	painpoint   PainPoint
	PainPointID int
	relevance   float64
}

//Create IssuePainPointsMap
func (issuepainpointmap *IssuePainPointsMap) Create() map[string]interface{} {
	GetDB().Create(issuepainpointmap)
	response := u.Message(true, "Issue has been created")
	response["issuepainpointmap"] = issuepainpointmap
	return response
}
