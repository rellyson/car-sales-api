package repositories

import "github.com/rellyson/car-sales-api/domain/entities"

type Entity interface {
	entities.Seller | entities.Car | entities.Ad
}

type GenericRepository[T Entity] interface {
	GetById(id string) (T, error)
	GetAll() ([]T, error)
	Create(T) (T, error)
	Update(T) (T, error)
}
