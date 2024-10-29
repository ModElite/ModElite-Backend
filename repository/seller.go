package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type sellerRepository struct {
	db *sqlx.DB
}

func NewSellerRepository(db *sqlx.DB) domain.SellerRepository {
	return &sellerRepository{
		db: db,
	}
}

func (r *sellerRepository) GetAll() (*[]domain.Seller, error) {
	sellers := make([]domain.Seller, 0)
	err := r.db.Select(&sellers, "SELECT * FROM seller")
	if err != nil {
		return nil, fmt.Errorf("error get all seller: %v", err)
	}
	return &sellers, nil
}

func (r *sellerRepository) GetByID(id string) (*domain.Seller, error) {
	var seller domain.Seller
	err := r.db.Get(&seller, "SELECT * FROM seller WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error get seller by id: %v", err)
	}
	return &seller, nil
}

func (r *sellerRepository) GetByOwnerID(ownerID string) (*[]domain.Seller, error) {
	sellers := make([]domain.Seller, 0)
	err := r.db.Select(&sellers, "SELECT * FROM seller WHERE owner_id = $1", ownerID)
	if err != nil {
		return nil, fmt.Errorf("error get seller by owner id: %v", err)
	}
	return &sellers, nil
}

func (r *sellerRepository) Create(seller *domain.Seller) error {
	_, err := r.db.NamedExec("INSERT INTO seller (id, name, description, logo_url, location, owner_id, is_verify) VALUES (:id, :name, :description, :logo_url, :location, :owner_id, :is_verify)", seller)
	if err != nil {
		return fmt.Errorf("error create seller: %v", err)
	}
	return nil
}

func (r *sellerRepository) Update(seller *domain.Seller) error {
	seller.UPDATED_AT = time.Now()
	_, err := r.db.NamedExec("UPDATE seller SET name = :name, description = :description, logo_url = :logo_url, location = :location, owner_id = :owner_id, is_verify = :is_verify, updated_at = :updated_at WHERE id = :id", seller)
	if err != nil {
		return fmt.Errorf("error update seller: %v", err)
	}
	return nil
}

func (r *sellerRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM seller WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error delete seller: %v", err)
	}
	return nil
}
