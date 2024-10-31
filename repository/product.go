package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetAll() (*[]domain.Product, error) {
	products := make([]domain.Product, 0)
	err := r.db.Select(&products, "SELECT * FROM product")
	if err != nil {
		return nil, fmt.Errorf("error getting product: %w", err)
	}
	return &products, nil
}

func (r *productRepository) GetByID(id string) (*domain.Product, error) {
	var product domain.Product
	err := r.db.Get(&product, "SELECT * FROM product WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error getting product: %w", err)
	}
	return &product, nil
}

func (r *productRepository) GetBySellerID(SellerID string) (*[]domain.Product, error) {
	products := make([]domain.Product, 0)
	err := r.db.Select(&products, "SELECT * FROM product WHERE seller_id = $1", SellerID)
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}
	return &products, nil
}

func (r *productRepository) Create(product *domain.Product) error {
	_, err := r.db.NamedExec("INSERT INTO product (id, seller_id, name, description, price, status) VALUES (:id, :seller_id, :name, :feature, :description, :price, :status)", product)
	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}
	return nil
}

func (r *productRepository) Update(product *domain.Product) error {
	product.UPDATED_AT = time.Now()
	_, err := r.db.NamedExec("UPDATE product SET name = :name, description = :description, price = :price, status = :status, updated_at = :updated_at WHERE id = :id", product)
	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}
	return nil
}

func (r *productRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM product WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}
	return nil
}
