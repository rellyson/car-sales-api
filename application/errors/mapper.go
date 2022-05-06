package errors

import (
	"errors"
	"net/http"
)

func MapError(w http.ResponseWriter, e error) error {
	return errors.New("")
}
