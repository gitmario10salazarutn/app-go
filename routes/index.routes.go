package routes

import (
	"encoding/json"
	"github.com/MarioUTN/api_rest_go/controller"
	"github.com/MarioUTN/api_rest_go/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HolaFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Mario, How are you?"))
}

func GetCompanys(w http.ResponseWriter, r *http.Request) {
	companys, err := controller.Get_AllCompanys()
	if err != nil {
		w.Write([]byte("Error to get Companys!"))
	} else {
		json.NewEncoder(w).Encode(companys)
	}
}

func GetCompany_ById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	id_empresa, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		w.Write([]byte("Error on Id"))
	}
	e, err := controller.Get_CompanyById(int(id_empresa))
	if err != nil {
		w.Write([]byte("Error on Get Company By Id"))
	} else {
		json.NewEncoder(w).Encode(e)
	}
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var e models.Empresa
	json.Unmarshal(requestBody, &e)
	id := mux.Vars(r)["id"]
	id_empresa, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		w.Write([]byte("Error on Id"))
	}
	err = controller.Update_Company(int(id_empresa), e)
	if err != nil {
		w.Write([]byte("Update failed!"))
	} else {
		w.Write([]byte("Update Company successfully!"))
	}
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	id_empresa, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		w.Write([]byte("Error on Id"))
		// We return, so we stop the function flow
	}
	err = controller.Delete_Company(int(id_empresa))
	if err != nil {
		w.Write([]byte("Delete failed!"))
	} else {
		w.Write([]byte("Delete Company successfully!"))
	}
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var e models.Empresa
	json.Unmarshal(requestBody, &e)
	err := controller.Create_Company(e)
	if err != nil {
		w.Write([]byte("Insert failed!"))
	} else {
		json.NewEncoder(w).Encode(&e)
	}
}
