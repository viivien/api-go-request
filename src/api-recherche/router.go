package main

import (
	"api-recherche/controllers"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/healthcheck").Name("healthcheck").HandlerFunc(controllers.HealthCheck)
	//router.Methods("GET").Path("/recherche").Name("recherche").HandlerFunc(controllers.RechercheOccurences)
	return router
}
