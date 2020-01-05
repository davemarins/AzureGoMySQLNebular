package controllers

import (
	"AzureGoMySQLNebular/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAdministrators(w http.ResponseWriter, r *http.Request) {
	admins := models.GetAllAdministrator()
	result, _ := json.Marshal(admins)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}

func GetAdministrator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminId := vars["userId"]
	id, err := strconv.ParseInt(adminId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	adminDetails, _ := models.GetAdministratorById(id)
	if adminDetails.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result, _ := json.Marshal(adminDetails)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}

func MakeAdministratorUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminId := vars["userId"]
	id, err := strconv.ParseInt(adminId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	adminTemp, db := models.GetUserById(id)
	if adminTemp.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if adminTemp.Role != string(models.ADMIN) {
		adminTemp.Role = string(models.ADMIN)
	}
	db.Save(&adminTemp)
	result, _ := json.Marshal(adminTemp)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}

func RevokeAdministratorStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminId := vars["userId"]
	id, err := strconv.ParseInt(adminId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	adminTemp, db := models.GetAdministratorById(id)
	if adminTemp.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if adminTemp.Role == string(models.ADMIN) {
		adminTemp.Role = string(models.USER)
	}
	db.Save(&adminTemp)
	result, _ := json.Marshal(adminTemp)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}
