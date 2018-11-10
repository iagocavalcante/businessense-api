package main

import (
	"businessense/app"
	"businessense/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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
	router.HandleFunc("/api/company/new", controllers.CreateCompany).Methods("POST")
	router.HandleFunc("/api/company/search", controllers.GetCompanies).Methods("GET")

	//APIs to handle PainPoints
	router.HandleFunc("/api/painpoint/new", controllers.CreatePainPoint).Methods("POST")
	router.HandleFunc("/api/painpoint/search", controllers.GetPainPoints).Methods("GET")
	router.HandleFunc("/api/painpoint/search/{search}", controllers.GetPainPoints).Methods("GET")

	//APIs to handle Issues
	router.HandleFunc("/api/issue/{id}", controllers.GetIssue).Methods("GET")
	router.HandleFunc("/api/issue/relevance", controllers.GetIssuesRelevance).Methods("POST")

	//APIs to handle Solutions
	router.HandleFunc("/api/solution/{id}", controllers.GetSolution).Methods("GET")
	router.HandleFunc("/api/solution/issue/{issueid}", controllers.GetSolutionsForIssue).Methods("GET")

	//APIs to handle Projects
	router.HandleFunc("/api/project/new", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/api/project/{id}", controllers.GetProject).Methods("GET")
	router.HandleFunc("/api/project/company/{id}", controllers.GetProjectByCompany).Methods("GET")

	//Deliver static files from public folder
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err := http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
