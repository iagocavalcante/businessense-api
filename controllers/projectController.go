package controllers

import (
	"businessense/models"
	u "businessense/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateProject HandlerFunc
var CreateProject = func(w http.ResponseWriter, r *http.Request) {

	project := &models.Project{}
	err := json.NewDecoder(r.Body).Decode(project) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := project.Create() //Create Project
	u.Respond(w, resp)
}

//GetProjectByCompany HandlerFunc
var GetProjectByCompany = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("id: ", id)
	data := models.GetProjectByCompany(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

//GetProject HandlerFunc
var GetProject = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("id: ", id)
	data := models.GetProject(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
