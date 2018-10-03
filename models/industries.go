package models

import (
	u "businessense/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Industry Type
type Industry struct {
	gorm.Model
	Name string `json:"name"`
}

//Create Industry
func (industry *Industry) Create() map[string]interface{} {
	GetDB().Create(industry)
	response := u.Message(true, "Industry has been created")
	response["industry"] = industry
	return response
}

//GetIndustries based on search string
func GetIndustries(search string) []*Industry {
	industries := make([]*Industry, 0)
	err := GetDB().Where("UPPER(name) LIKE UPPER(?)", "%"+search+"%").Find(&industries).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return industries
}
