package main

import (
	"github.com/rellyson/car-sales-api/infra/http"
	"github.com/rellyson/car-sales-api/infra/persistence"
)

func main() {
	//create and instantiate database
	persistence.CreateDBConnection(persistence.DatabaseConfig{
		Driver:   "postgres",
		User:     "cs_admin",
		Password: "custompassword",
		Host:     "0.0.0.0",
		Port:     5432,
		DbName:   "car_sales_db",
		SSLMode:  "disabled",
	})

	//check and exec pending migrations
	persistence.ExecMigrations(persistence.GetDBConnection(), "car_sales_db")

	//opens http server
	http.CreateHttpServer(":3000")
}
