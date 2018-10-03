package main

import (
	"businessense/app"
	"businessense/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//Apis to handle user creation and loging
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	//Apis to handle industry creation and listing
	router.HandleFunc("/api/industry/new", controllers.CreateIndustry).Methods("POST")
	router.HandleFunc("/api/industry/search", controllers.GetIndustries).Methods("GET")
	router.HandleFunc("/api/industry/search/{search}", controllers.GetIndustries).Methods("GET")

	//APIs to handle company creation and listing
	//router.HandleFunc("/api/company/{id}", controllers.GetCompny)
	router.HandleFunc("/api/company/new", controllers.CreateCompany).Methods("POST")
	router.HandleFunc("/api/companies", controllers.GetCompanies).Methods("GET")

	//Deliver static files from public folder
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
