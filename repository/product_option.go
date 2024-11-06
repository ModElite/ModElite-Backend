package repository

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type productOptionRepository struct {
	db *sqlx.DB
}

func NewProductOptionRepository(db *sqlx.DB) domain.ProductOptionRepository {
	return &productOptionRepository{
		db: db,
	}
}

func (r *productOptionRepository) GetAll() (*[]domain.ProductOption, error) {
	productOptions := make([]domain.ProductOption, 0)
	err := r.db.Select(&productOptions, "SELECT * FROM product_option")
	if err != nil {
		return nil, err
	}
	return &productOptions, nil
}

func (r *productOptionRepository) GetByID(id string) (*domain.ProductOption, error) {
	var productOption domain.ProductOption
	err := r.db.Get(&productOption, "SELECT * FROM product_option WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &productOption, nil
}

func (r *productOptionRepository) GetByProductID(productID string) (*[]domain.ProductOption, error) {
	productOptions := make([]domain.ProductOption, 0)
	err := r.db.Select(&productOptions, "SELECT * FROM product_option WHERE product_id = $1", productID)
	if err != nil {
		return nil, err
	}
	return &productOptions, nil
}

func (r *productOptionRepository) Create(productOption *domain.ProductOption) error {
	_, err := r.db.NamedExec("INSERT INTO product_option (id, product_id, label, image_url, price) VALUES (:id, :product_id, :label, :image_url, :price)", productOption)
	if err != nil {
		return err
	}
	return nil
}

func (r *productOptionRepository) Update(productOption *domain.ProductOption) error {
	_, err := r.db.NamedExec("UPDATE product_option SET label = :label, price = :price WHERE id = :id", productOption)
	if err != nil {
		return err
	}
	return nil
}
