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
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *addressRepository) GetByUserId(userId string) (*[]domain.Address, error) {
	addresses := make([]domain.Address, 0)
	err := r.db.Select(&addresses, `SELECT * FROM address WHERE user_id = $1 ORDER BY "default" DESC`, userId)
	if err != nil {
		return nil, err
	}

	return &addresses, nil
}

func (r *addressRepository) Create(address *domain.Address) (int, error) {
	var id int
	query := `
		INSERT INTO address (
			user_id, first_name, last_name, email, phone, label, "default", address, sub_district, district, province, zip_code
		) VALUES (
			:user_id, :first_name, :last_name, :email, :phone, :label, :default, :address, :sub_district, :district, :province, :zip_code
		) RETURNING id`

	rows, err := r.db.NamedQuery(query, address)
	if err != nil {
		return 0, fmt.Errorf("failed to create address: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("failed to retrieve id: %w", err)
		}
	}

	return id, nil
}

func (r *addressRepository) Update(address *domain.Address) error {
	_, err := r.db.NamedExec(`UPDATE address SET first_name = :first_name, last_name = :last_name, email = :email, phone = :phone, label = :label, "default" = :default, address = :address, sub_district = :sub_district, district = :district, province = :province, zip_code = :zip_code, updated_at = now() WHERE id = :id`, address)
	if err != nil {
		return fmt.Errorf("failed to update address: %w", err)
	}

	return nil
}

func (r *addressRepository) UpdateDefaultByUserId(userId string, id int) error {
	tx := r.db.MustBegin()
	_, err := tx.Exec(`UPDATE address SET "default" = false, updated_at = NOW() WHERE user_id = $1`, userId)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return fmt.Errorf("failed to update address: %w", errRollback)
		}
		return fmt.Errorf("failed to update address: %w", err)
	}
	_, err = tx.Exec(`UPDATE address SET "default" = true, updated_at = NOW() WHERE id = $1`, id)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return fmt.Errorf("failed to update address: %w", errRollback)
		}
		return fmt.Errorf("failed to update address: %w", err)
	}

	if err = tx.Commit(); err != nil {
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
