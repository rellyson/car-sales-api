package mappers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rellyson/car-sales-api/domain/entities"
)

type SellerMap struct {
	BaseMap[entities.Seller]
}

func (m *SellerMap) ToDomain(data map[string]any) (entities.Seller, error) {
	id, _ := uuid.Parse(fmt.Sprint(data["id"]))
	c, _ := time.Parse(time.ANSIC, fmt.Sprint(data["created_at"]))
	u, _ := time.Parse(time.ANSIC, fmt.Sprint(data["updated_at"]))
	seller := entities.Seller{
		BaseEntity: entities.BaseEntity{
			ID:        id,
			CreatedAt: c,
			UpdatedAt: u,
		},
		FullName: fmt.Sprint(data["full_name"]),
		Email:    fmt.Sprint(data["email"]),
		Password: fmt.Sprint(data["password"]),
	}

	return seller, nil
}
