package repositories

import "github.com/rellyson/car-sales-api/domain/entities"

type Entity interface {
	entities.Seller | entities.Car | entities.Ad
}

type GenericRepository[T Entity] interface {
	GetById(id string) T
	GetAll() []T
	Create(T) T
	Update(T) T
}
