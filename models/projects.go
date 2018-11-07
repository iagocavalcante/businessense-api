package models

import (
	u "businessense/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Project Type
type Project struct {
	gorm.Model
	Name       string `json:"name"`
	Company    Company
	CompanyID  int
	PainPoints []PainPoint `gorm:"many2many:project_painpoint;"`
	Issues     []Issue     `gorm:"many2many:project_issue;"`
	Solutions  []Solution  `gorm:"many2many:project_solution;"`
}

//Create Project
func (project *Project) Create() map[string]interface{} {
	GetDB().Create(project)
	response := u.Message(true, "Project has been created")
	response["project"] = project
	return response
}

//GetSolution based on search string
func GetProjectByCompany(company int) []*Project {
	projects := make([]*Project, 0)
	err := GetDB().Where("company_id = ?", company).Find(&projects).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return projects
}
