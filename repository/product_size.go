package repository

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type productSizeRepository struct {
	db *sqlx.DB
}

func NewProductSizeRepository(db *sqlx.DB) domain.ProductSizeRepository {
	return &productSizeRepository{
		db: db,
	}
}

func (r *productSizeRepository) GetAll() (*[]domain.ProductSize, error) {
	productSizes := make([]domain.ProductSize, 0)
	err := r.db.Select(&productSizes, "SELECT * FROM product_size")
	if err != nil {
		return nil, fmt.Errorf("error cannot get product sizes: %w", err)
	}
	return &productSizes, nil
}

func (r *productSizeRepository) GetByID(id string) (*domain.ProductSize, error) {
	var productSize domain.ProductSize
	err := r.db.Get(&productSize, "SELECT * FROM product_size WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error cannot get product size: %w", err)
	}
	return &productSize, nil
}

func (r *productSizeRepository) GetByProductOptionID(productOptionID string) (*[]domain.ProductSize, error) {
	productSizes := make([]domain.ProductSize, 0)
	err := r.db.Select(&productSizes, "SELECT * FROM product_size WHERE product_option_id = $1", productOptionID)
	if err != nil {
		return nil, fmt.Errorf("error cannot get product sizes: %w", err)
	}
	return &productSizes, nil
}

func (r *productSizeRepository) Create(productSize *domain.ProductSize) error {
	_, err := r.db.NamedExec("INSERT INTO product_size (id, product_option_id, size_id, quantity) VALUES (:id, :product_option_id, :size_id, :quantity)", productSize)
	if err != nil {
		return fmt.Errorf("error cannot create product size: %w", err)
	}
	return nil
}

func (r *productSizeRepository) Update(productSize *domain.ProductSize) error {
	_, err := r.db.NamedExec("UPDATE product_size SET quantity = :quantity WHERE id = :id", productSize)
	if err != nil {
		return fmt.Errorf("error cannot update product size: %w", err)
	}
	return nil
}
