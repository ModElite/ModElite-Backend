package repository

import (
	"database/sql"
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type productTagsRepository struct {
	db *sqlx.DB
}

func NewProductTagsRepository(db *sqlx.DB) domain.ProductTagsRepository {
	return &productTagsRepository{
		db: db,
	}
}

func (r *productTagsRepository) GetAll() (*[]domain.ProductTag, error) {
	productTags := make([]domain.ProductTag, 0)
	err := r.db.Select(&productTags, "SELECT * FROM product_tags")
	if err != nil {
		return nil, fmt.Errorf("error while getting all product tags: %w", err)
	}
	return &productTags, nil
}

func (r *productTagsRepository) GetByProductID(productID string) (*[]domain.ProductTag, error) {
	productTags := make([]domain.ProductTag, 0)
	err := r.db.Select(&productTags, "SELECT * FROM product_tags WHERE product_id = $1", productID)
	if err != nil {
		return nil, fmt.Errorf("error while getting product tag by product id: %w", err)
	}
	return &productTags, nil
}

func (r *productTagsRepository) GetByTagID(tagID int) (*[]domain.ProductTag, error) {
	productTags := make([]domain.ProductTag, 0)
	err := r.db.Select(&productTags, "SELECT * FROM product_tags WHERE tag_id = $1", tagID)
	if err != sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error while getting product tag by tag id: %w", err)
	}
	return &productTags, nil
}

func (r *productTagsRepository) Create(productID string, tagID int) error {
	_, err := r.db.Exec("INSERT INTO product_tags(product_id, tag_id) VALUES($1, $2)", productID, tagID)
	if err != nil {
		return fmt.Errorf("error while creating product tag: %w", err)
	}
	return nil
}

func (r *productTagsRepository) Delete(productID string, tagID int) error {
	_, err := r.db.Exec("DELETE FROM product_tags WHERE product_id = $1 AND tag_id = $2", productID, tagID)
	if err != nil {
		return fmt.Errorf("error while deleting product tag: %w", err)
	}
	return nil
}
