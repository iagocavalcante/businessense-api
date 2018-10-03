package models

import (
	u "businessense/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Company from Accounts
type Company struct {
	gorm.Model
	Name       string  `json:"name"`
	Revenue    float64 `json:"revenue"`
	Account    Account
	AccountID  int `json:"accountid"`
	Industry   Industry
	IndustryID int `json:"industryid"`
}

//Create Company
func (company *Company) Create() map[string]interface{} {
	GetDB().Create(company)
	response := u.Message(true, "Company has been created")
	response["company"] = company
	return response
}

//GetCompanies based on accountid
func GetCompanies(accountid uint) []*Company {
	companies := make([]*Company, 0)
	err := GetDB().Set("gorm:auto_preload", true).Where("account_id = ?", accountid).Find(&companies).Error

	for i := range companies {
		companies[i].Account.Password = ""
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return companies
}
