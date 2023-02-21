package main

import (
	"github.com/MarioUTN/api_rest_go/database"
	"github.com/MarioUTN/api_rest_go/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	database.DatabaseConnection()
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HolaFunction)
	r.HandleFunc("/get_empresas", routes.GetCompanys).Methods("GET")
	r.HandleFunc("/get_empresabyid/{id}", routes.GetCompany_ById).Methods("GET")
	r.HandleFunc("/update_empresa/{id}", routes.UpdateCompany).Methods("PUT")
	r.HandleFunc("/delete_empresa/{id}", routes.DeleteCompany).Methods("DELETE")
	r.HandleFunc("/create_empresa", routes.CreateCompany).Methods("POST")

	http.ListenAndServe(":3000", r)
}
