package usecases

import (
	"github.com/rellyson/car-sales-api/domain/dtos"
	"github.com/rellyson/car-sales-api/domain/entities"
	"github.com/rellyson/car-sales-api/domain/repositories"
)

type CreateSellerUseCase interface {
	Handle(data dtos.CreateSellerDTO) (any, error)
}

type createSellerUseCase struct{}

var repo repositories.GenericRepository[entities.Seller]

func NewCreateSellerUseCase(r repositories.GenericRepository[entities.Seller]) CreateSellerUseCase {
	repo = r
	return &createSellerUseCase{}
}

func (*createSellerUseCase) Handle(data dtos.CreateSellerDTO) (any, error) {
	seller := entities.Seller{
		FullName: data.FullName,
		Email:    data.Email,
	}

	err := seller.HashPassword(data.Password)

	if err != nil {
		return nil, err
	}

	res, err := repo.Create(seller)

	if err != nil {
		return nil, err
	}

	return res, nil
}
