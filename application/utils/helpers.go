package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ParseJSONBody(r *http.Request, s interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		return err
	}

	return nil
}
