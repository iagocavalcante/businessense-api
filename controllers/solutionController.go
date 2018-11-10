package controllers

import (
	"businessense/models"
	u "businessense/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetSolutionsForIssue HandlerFunc
var GetSolutionsForIssue = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("id: ", id)
	dataSolutions := models.GetSolutionsForIssue(id)
	constraintSolutions := models.GetConstraintsForIssue(id)
	resp := u.Message(true, "success")
	resp["solutions"] = dataSolutions
	resp["constraints"] = constraintSolutions
	u.Respond(w, resp)
}

//GetSolution HandlerFunc
var GetSolution = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("id: ", id)
	data := models.GetSolution(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
