package http

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/rellyson/car-sales-api/application/controllers"
	"github.com/rellyson/car-sales-api/application/utils"
	"github.com/rellyson/car-sales-api/infra/http/middlewares"
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
	logger := utils.NewLogger()

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()

		logger.Info(fmt.Sprintf("[RouteMapper] - Route %v %v mapped.", methods, pathTemplate))
		return nil
	})

	if err != nil {
		logger.Error(err.Error())
	}
}
