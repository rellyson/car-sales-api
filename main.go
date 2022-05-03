package main

import "github.com/rellyson/car-sales-api/infra/http"

func main() {
	http.CreateHttpServer(":3000")
}
