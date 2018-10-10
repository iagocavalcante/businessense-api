package models

import (
	u "businessense/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//PainPoint Type
type PainPoint struct {
	gorm.Model
	Name string `json:"name"`
}

//Create PainPoint
func (painpoint *PainPoint) Create() map[string]interface{} {
	GetDB().Create(painpoint)
	response := u.Message(true, "Pain Point has been created")
	response["painpoint"] = painpoint
	return response
}

//GetPainpoints based on search string
func GetPainpoints(search string) []*PainPoint {
	painpoints := make([]*PainPoint, 0)
	err := GetDB().Where("UPPER(name) LIKE UPPER(?)", "%"+search+"%").Find(&painpoints).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return painpoints
}
