package usecases

import (
	"github.com/rellyson/car-sales-api/domain/entities"
	"github.com/rellyson/car-sales-api/domain/repositories"
)

type GetSellerByIdUseCase interface {
	Handle(id string) (any, error)
}

type getSellerByIdUseCase struct {
	repo repositories.GenericRepository[entities.Seller]
}

func NewGetSellerByIdUseCase(r repositories.GenericRepository[entities.Seller]) GetSellerByIdUseCase {
	return &getSellerByIdUseCase{
		repo: r,
	}
}

func (us *getSellerByIdUseCase) Handle(id string) (any, error) {
	res, err := us.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
