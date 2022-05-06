package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rellyson/car-sales-api/application/errors"
	"github.com/rellyson/car-sales-api/application/utils"
	"github.com/rellyson/car-sales-api/domain/dtos"
	usecases "github.com/rellyson/car-sales-api/domain/use_cases"
)

type sellerController struct{}

var (
	createSellerUs usecases.CreateSellerUseCase
)

type SellerController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func NewSellerController(createUs usecases.CreateSellerUseCase) SellerController {
	createSellerUs = createUs

	return &sellerController{}
}

func (*sellerController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (*sellerController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (*sellerController) Create(w http.ResponseWriter, r *http.Request) {
	dto := &dtos.CreateSellerDTO{}
	if err := utils.ParseJSONBody(r.Body, dto); err != nil {
		errors.MapError(w, err)
		return
	}

	if err := dto.Validate(); err != nil {
		errors.MapError(w, err)
		return
	}

	res, err := createSellerUs.Handle(dto)

	if err != nil {
		errors.MapError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (*sellerController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
