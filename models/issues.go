package models

import (
	u "businessense/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Issue Type
type Issue struct {
	gorm.Model
	Name string `json:"name"`
}

//Create Issue
func (issue *Issue) Create() map[string]interface{} {
	GetDB().Create(issue)
	response := u.Message(true, "Issue has been created")
	response["issue"] = issue
	return response
}

//GetIssues based on search string
func GetIssues(search string) []*Issue {
	issue := make([]*Issue, 0)
	err := GetDB().Where("UPPER(name) LIKE UPPER(?)", "%"+search+"%").Find(&issue).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return issue
}
