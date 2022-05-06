package utils

import (
	"encoding/json"
	"io"
)

func ParseJSONBody(r io.ReadCloser, s interface{}) error {
	body, err := io.ReadAll(r)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), s)

	return nil
}
