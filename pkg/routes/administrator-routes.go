package routes

import (
	"AzureGoMySQLNebular/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterAdministratorRoutes = func(router *mux.Router) {
	router.HandleFunc("/administrator/", controllers.GetAdministrators).Methods("GET")
	router.HandleFunc("/administrator/{userId}", controllers.GetAdministrator).Methods("GET")
	router.HandleFunc("/administrator/{userId}", controllers.MakeAdministratorUser).Methods("PUT")
	router.HandleFunc("/administrator/{userId}", controllers.RevokeAdministratorStatus).Methods("DELETE")
}
