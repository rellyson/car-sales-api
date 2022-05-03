package controllers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rellyson/car-sales-api/application/controllers"
	"github.com/stretchr/testify/assert"
)

func TestHCStatus(t *testing.T) {
	c := controllers.NewHealthCheckController()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	c.Status(rr, req)
	res := rr.Result()

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Check the response body is what we expect.
	resBody, _ := io.ReadAll(res.Body)
	assert.JSONEq(t, `{"status":200, "message": "Ok"}`, string(resBody))
}
