package models

import (
	u "businessense/utils"

	"github.com/jinzhu/gorm"
)

//Solution Type
type Solution struct {
	gorm.Model
	Name           string `json:"name"`
	Description    string `json:"description"`
	Implementation string `json:"implementation"`
	SkillsTools    string `json:"skillstools"`
}

//Create Solution
func (solution *Solution) Create() map[string]interface{} {
	GetDB().Create(solution)
	response := u.Message(true, "Solution has been created")
	response["solution"] = solution
	return response
}

//GetSolutionsForIssue based on issue id
func GetSolutionsForIssue(issueid int) []*Solution {
	solutions := make([]*Solution, 0)

	GetDB().Find(&solutions).Where("exists (select 1 from constraint_issues_solution_maps where solutions.id = solution_id and issue_id = ?)", issueid)
	return solutions
}

//GetSolution based on search string
func GetSolution(id int) Solution {
	var solution Solution
	GetDB().First(&solution, id)
	return solution
}
