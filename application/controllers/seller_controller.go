package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rellyson/car-sales-api/application/errors"
	"github.com/rellyson/car-sales-api/application/utils"
	"github.com/rellyson/car-sales-api/domain/dtos"
	usecases "github.com/rellyson/car-sales-api/domain/use_cases"
)

type sellerController struct {
	createSellerUs usecases.CreateSellerUseCase
	getSellerById  usecases.GetSellerByIdUseCase
}

type SellerController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func NewSellerController(createUs usecases.CreateSellerUseCase, getSellerByIdUs usecases.GetSellerByIdUseCase) SellerController {
	return &sellerController{
		createSellerUs: createUs,
		getSellerById:  getSellerByIdUs,
	}
}

func (c *sellerController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		return
	}

	res, err := c.getSellerById.Handle(id)

	if err != nil {
		errors.MapError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (*sellerController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (c *sellerController) Create(w http.ResponseWriter, r *http.Request) {
	dto := &dtos.CreateSellerDTO{}
	if err := utils.ParseJSONBody(r.Body, dto); err != nil {
		errors.MapError(w, err)
		return
	}

	if err := dto.Validate(); err != nil {
		errors.MapError(w, err)
		return
	}

	res, err := c.createSellerUs.Handle(dto)

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
