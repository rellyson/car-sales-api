package controllers

import (
	"encoding/json"
	"net/http"
)

type healthCheckController struct{}

type HealthCheckController interface {
	Status(w http.ResponseWriter, r *http.Request)
}

func NewHealthCheckController() HealthCheckController {
	return &healthCheckController{}
}

func (*healthCheckController) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  http.StatusOK,
		Message: "Ok",
	}

	json.NewEncoder(w).Encode(res)
}
