package http

import (
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/rellyson/car-sales-api/application/controllers"
	"github.com/rellyson/car-sales-api/infra/http/middlewares"
	"github.com/rellyson/car-sales-api/infra/utils"
)

var (
	healthCheckController controllers.HealthCheckController = controllers.NewHealthCheckController()
)

func SetRoutes() *mux.Router {
	r := mux.NewRouter()

	//set routes
	r.HandleFunc("/healthcheck", healthCheckController.Status).Methods("GET")

	//router mapper and set logger
	mapRoutes(r)
	r.Use(middlewares.LoggingMiddleware)

	return r
}

func mapRoutes(r *mux.Router) {
	colorizedLog := utils.NewColorizedLogger(color.FgGreen)

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()

		colorizedLog.Printf("[RouteMapper] - Route %v %v mapped. \n", methods, pathTemplate)
		return nil
	})

	if err != nil {
		colorizedLog.Add(color.FgRed).Println(err)
	}
}
