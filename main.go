package main

import (
	"github.com/rellyson/car-sales-api/infra/http"
	"github.com/rellyson/car-sales-api/infra/persistence"
)

func main() {
	persistence.CreateDBConnection(persistence.DatabaseConfig{
		Driver:   "postgres",
		User:     "cs_admin",
		Password: "custompassword",
		Host:     "0.0.0.0",
		Port:     5432,
		DbName:   "car_sales_db",
		SSLMode:  "disabled",
	})

	http.CreateHttpServer(":3000")
}
