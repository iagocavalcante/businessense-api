package controllers

import (
	"businessense/models"
	u "businessense/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//CreateCompany HandlerFunc
var CreateCompany = func(w http.ResponseWriter, r *http.Request) {

	company := &models.Company{}
	err := json.NewDecoder(r.Body).Decode(company) //decode the request body into struct and failed if any error occur

	fmt.Println(company.Name)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := company.Create() //Create account
	u.Respond(w, resp)
}

//GetCompanies HandlerFunc
var GetCompanies = func(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	accountid := r.Context().Value("user") //params["accountid"]
	fmt.Println("accountid: ", accountid)
	data := models.GetCompanies(accountid.(uint))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
