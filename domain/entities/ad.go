package entities

import "time"

type Ad struct {
	BaseEntity
	Value       float32   `json:"value"`
	IsAvailable bool      `json:"is_available"`
	SoldIn      time.Time `json:"sold_in"`
	Seller      Seller
	Car         Car
}
