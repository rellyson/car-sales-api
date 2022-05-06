package dtos

import "github.com/rellyson/car-sales-api/application/utils"

type CreateSellerDTO struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (dto *CreateSellerDTO) Validate() error {
	err := utils.ValidateStruct(dto)

	if err != nil {
		return err
	}

	return nil
}
