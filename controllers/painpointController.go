package controllers

import (
	"businessense/models"
	u "businessense/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//CreatePainPoint HandlerFunc
var CreatePainPoint = func(w http.ResponseWriter, r *http.Request) {

	painpoint := &models.PainPoint{}
	err := json.NewDecoder(r.Body).Decode(painpoint) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := painpoint.Create() //Create PainPoint
	u.Respond(w, resp)
}

//GetPainPoints HandlerFunc
var GetPainPoints = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	search := params["search"]
	fmt.Println("search: " + search)
	data := models.GetPainpoints(search)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
