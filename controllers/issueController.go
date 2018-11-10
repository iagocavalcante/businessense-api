package controllers

import (
	"businessense/models"
	u "businessense/utils"
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/gorilla/mux"
)

//GetIssuesRelevance HandlerFunc
var GetIssuesRelevance = func(w http.ResponseWriter, r *http.Request) {

	painpoints := make([]string, 0)
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&painpoints) //decode the request body into struct and failed if any error occur

	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := models.GetIssuesRelevance(painpoints) //Get issues relevance based on painpoints
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

//GetIssue based on Id
var GetIssue = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("id: ", id)
	data := models.GetIssue(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
