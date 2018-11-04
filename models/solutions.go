package models

import (
	u "businessense/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Constraint Type
type Solution struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

//Create Constraint
func (solution *Solution) Create() map[string]interface{} {
	GetDB().Create(solution)
	response := u.Message(true, "Solution has been created")
	response["solution"] = solution
	return response
}

//GetSolution based on search string
func GetSolution(search string) []*Solution {
	solutions := make([]*Solution, 0)
	err := GetDB().Where("UPPER(name) LIKE UPPER(?)", "%"+search+"%").Find(&solutions).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return solutions
}
