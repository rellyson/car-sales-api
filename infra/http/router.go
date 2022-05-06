package http

import (
	"database/sql"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rellyson/car-sales-api/application/controllers"
	"github.com/rellyson/car-sales-api/application/repositories"
	"github.com/rellyson/car-sales-api/application/utils"
	"github.com/rellyson/car-sales-api/domain/entities"
	domainrepo "github.com/rellyson/car-sales-api/domain/repositories"
	usecases "github.com/rellyson/car-sales-api/domain/use_cases"
	"github.com/rellyson/car-sales-api/infra/http/middlewares"
	"github.com/rellyson/car-sales-api/infra/persistence"
)

var (
	//db
	db *sql.DB = persistence.GetDBConnection()

	//repositories
	sellerRepo domainrepo.GenericRepository[entities.Seller] = repositories.NewSellerRepositoryImp(db)

	//usecases
	createSellerUseCase usecases.CreateSellerUseCase = usecases.NewCreateSellerUseCase(sellerRepo)

	//controllers
	healthCheckController controllers.HealthCheckController = controllers.NewHealthCheckController()
	sellerController      controllers.SellerController      = controllers.NewSellerController(createSellerUseCase)
)

func SetRoutes() *mux.Router {
	r := mux.NewRouter()

	//set routes
	r.HandleFunc("/healthcheck", healthCheckController.Status).Methods("GET")
	r.HandleFunc("/sellers/{id}", sellerController.GetById).Methods("GET")
	r.HandleFunc("/sellers", sellerController.GetAll).Methods("GET")
	r.HandleFunc("/sellers", sellerController.Create).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/sellers", sellerController.Update).Methods("PUT").Headers("Content-Type", "application/json")

	configureRouter(r)

	return r
}

func mapRoutes(r *mux.Router) {
	logger := utils.NewLogger()

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		handlerPath := strings.Split(runtime.FuncForPC(reflect.ValueOf(route.GetHandler()).Pointer()).Name(), ".")
		handlerController := handlerPath[len(handlerPath)-2]
		handlerFunc := strings.Split(handlerPath[len(handlerPath)-1], "-")[0]
		routeHandler := fmt.Sprintf("%s.%s", handlerController, handlerFunc)

		logger.Info(fmt.Sprintf("[RouteMapper] - %v %v -> %v mapped.", methods, pathTemplate, routeHandler))
		return nil
	})

	if err != nil {
		logger.Error(err.Error())
	}
}

func configureRouter(r *mux.Router) {
	//route mapper and set logger
	mapRoutes(r)
	r.Use(middlewares.LoggingMiddleware)
}
