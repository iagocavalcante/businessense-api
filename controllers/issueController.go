package controllers

import (
	"businessense/models"
	u "businessense/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GetIssues HandlerFunc
var GetIssues = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	painPoints := params["painpoints"]
	fmt.Println("painPoints: " + painPoints)
	data := models.GetIssues(painPoints)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
