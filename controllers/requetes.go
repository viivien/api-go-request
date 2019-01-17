package controllers

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	status string
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	statusOK := Status{
		status: "OK",
	}

	json.NewEncoder(w).Encode(statusOK)
}
