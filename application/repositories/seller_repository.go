package repositories

import (
	"database/sql"
	"errors"

	"github.com/rellyson/car-sales-api/domain/entities"
	"github.com/rellyson/car-sales-api/domain/repositories"
)

type sellerRepositoryImp struct{}

var db *sql.DB

func NewSellerRepositoryImp(database *sql.DB) repositories.GenericRepository[entities.Seller] {
	db = database
	return &sellerRepositoryImp{}
}

func (*sellerRepositoryImp) GetById(id string) (entities.Seller, error) {
	seller := entities.Seller{}
	rows, err := db.Query("SELECT id, full_name, email, password, created_at, updated_at FROM sellers WHERE id = $1", id)

	if err != nil {
		return seller, errors.New(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			seller.ID,
			seller.FullName,
			seller.Email,
			seller.Password,
			seller.CreatedAt,
			seller.UpdatedAt,
		)
	}

	if rows.Err(); err != nil {
		return seller, errors.New(err.Error())
	}

	return seller, nil
}

func (*sellerRepositoryImp) GetAll() ([]entities.Seller, error) {
	sellers := []entities.Seller{}
	rows, err := db.Query("SELECT id, full_name, email, created_at, updated_at FROM sellers")

	if err != nil {
		return sellers, errors.New(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rowSeller := entities.Seller{}
		rows.Scan(
			rowSeller.ID,
			rowSeller.FullName,
			rowSeller.Email,
			rowSeller.CreatedAt,
			rowSeller.UpdatedAt,
		)

		sellers = append(sellers, rowSeller)
	}

	if rows.Err(); err != nil {
		return sellers, errors.New(err.Error())
	}

	return sellers, nil
}

func (*sellerRepositoryImp) Create(s entities.Seller) (entities.Seller, error) {
	seller := entities.Seller{}
	stmt, err := db.Prepare("INSERT INTO sellers(full_name, email, password) VALUES($1, $2, $3) RETURNING id, full_name, email, created_at, updated_at")

	if err != nil {
		return seller, errors.New(err.Error())
	}

	err = stmt.QueryRow(s.FullName, s.Email, s.Password).Scan(
		seller.ID,
		seller.FullName,
		seller.Email,
		seller.CreatedAt,
		seller.UpdatedAt,
	)

	if err != nil {
		return seller, errors.New(err.Error())
	}

	return seller, nil
}

func (*sellerRepositoryImp) Update(s entities.Seller) (entities.Seller, error) {
	seller := entities.Seller{}
	stmt, err := db.Prepare(`
	UPDATE sellers SET full_name = $1, email = $2, password = $3, updated_at = NOW()
	WHERE id = $1
	RETURNING id, full_name, email, created_at, updated_at
	`)

	if err != nil {
		return seller, errors.New(err.Error())
	}

	err = stmt.QueryRow(s.FullName, s.Email, s.Password).Scan(
		seller.ID,
		seller.FullName,
		seller.Email,
		seller.CreatedAt,
		seller.UpdatedAt,
	)

	if err != nil {
		return seller, errors.New(err.Error())
	}

	return seller, nil
}
