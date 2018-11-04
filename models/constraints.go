package models

import (
	u "businessense/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Constraint Type
type Constraint struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

//Create Constraint
func (constraint *Constraint) Create() map[string]interface{} {
	GetDB().Create(constraint)
	response := u.Message(true, "Constraint has been created")
	response["constraint"] = constraint
	return response
}

//GetConstraint based on search string
func GetConstraint(search string) []*Constraint {
	constraints := make([]*Constraint, 0)
	err := GetDB().Where("UPPER(name) LIKE UPPER(?)", "%"+search+"%").Find(&constraints).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return constraints
}
