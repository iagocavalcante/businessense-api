package controllers

import (
	"businessense/models"
	u "businessense/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//CreateIndustry HandlerFunc
var CreateIndustry = func(w http.ResponseWriter, r *http.Request) {

	industry := &models.Industry{}
	err := json.NewDecoder(r.Body).Decode(industry) //decode the request body into struct and failed if any error occur

	fmt.Println(industry.Name)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := industry.Create() //Create account
	u.Respond(w, resp)
}

//GetIndustries HandlerFunc
var GetIndustries = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	search := params["search"]
	fmt.Println("search: " + search)
	data := models.GetIndustries(search)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
