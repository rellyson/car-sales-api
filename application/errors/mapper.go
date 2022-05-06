package errors

import (
	"encoding/json"
	"net/http"
)

type HttpJSONError struct {
	Status  int    `json:"status"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func MapError(w http.ResponseWriter, e error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	body := &HttpJSONError{
		Status:  http.StatusBadRequest,
		Type:    "",
		Message: e.Error(),
	}
	json.NewEncoder(w).Encode(body)
}
