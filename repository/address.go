package repository

import (
	"database/sql"
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type addressRepository struct {
	db *sqlx.DB
}

func NewAddressRepository(db *sqlx.DB) domain.AddressRepository {
	return &addressRepository{
		db: db,
	}
}

func (r *addressRepository) GetAll() (*[]domain.Address, error) {
	addresses := make([]domain.Address, 0)
	err := r.db.Select(&addresses, "SELECT * FROM address")
	if err != nil {
		return nil, err
	}

	return &addresses, nil
}

func (r *addressRepository) GetById(id int) (*domain.Address, error) {
	address := domain.Address{}
	err := r.db.Get(&address, "SELECT * FROM address WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("address with id %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *addressRepository) GetByUserId(userId string) (*[]domain.Address, error) {
	addresses := make([]domain.Address, 0)
	err := r.db.Select(&addresses, "SELECT * FROM address WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	return &addresses, nil
}

func (r *addressRepository) Create(address *domain.Address) error {
	_, err := r.db.NamedExec("INSERT INTO address (user_id, first_name, last_name, company, street_address, state, country, zip_code, email, phone, type)"+
		" VALUES (:user_id, :first_name, :last_name, :company, :street_address, :state, :country, :zip_code, :email, :phone, :type)", address)
	if err != nil {
		return fmt.Errorf("failed to create address: %w", err)
	}

	return nil
}

func (r *addressRepository) Update(address *domain.Address) error {
	_, err := r.db.NamedExec("UPDATE address SET user_id = :user_id, first_name = :first_name, last_name = :last_name, company = :company, street_address = :street_address, state = :state, country = :country, zip_code = :zip_code, email = :email, phone = :phone, type = :type WHERE id = :id", address)
	if err != nil {
		return fmt.Errorf("failed to update address: %w", err)
	}

	return nil
}

func (r *addressRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM address WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete address: %w", err)
	}

	return nil
}
