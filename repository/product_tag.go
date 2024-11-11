package repository

import (
	"database/sql"
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type productTagRepository struct {
	db *sqlx.DB
}

func NewProductTagRepository(db *sqlx.DB) domain.ProductTagRepository {
	return &productTagRepository{
		db: db,
	}
}

func (r *productTagRepository) GetAll() (*[]domain.ProductTag, error) {
	productTags := make([]domain.ProductTag, 0)
	err := r.db.Select(&productTags, "SELECT * FROM product_tag")
	if err != nil {
		return nil, fmt.Errorf("error while getting all product tags: %w", err)
	}
	return &productTags, nil
}

func (r *productTagRepository) GetByProductID(productID string) (*[]domain.ProductTag, error) {
	productTags := make([]domain.ProductTag, 0)
	err := r.db.Select(&productTags, "SELECT * FROM product_tag WHERE product_id = $1", productID)
	if err != nil {
		return nil, fmt.Errorf("error while getting product tag by product id: %w", err)
	}
	return &productTags, nil
}

func (r *productTagRepository) GetByTagID(tagID int) (*[]domain.ProductTag, error) {
	productTags := make([]domain.ProductTag, 0)
	err := r.db.Select(&productTags, "SELECT * FROM product_tag WHERE tag_id = $1", tagID)
	if err != sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error while getting product tag by tag id: %w", err)
	}
	return &productTags, nil
}

func (r *productTagRepository) Create(productID string, tagID int) error {
	_, err := r.db.Exec("INSERT INTO product_tag(product_id, tag_id) VALUES($1, $2)", productID, tagID)
	if err != nil {
		return fmt.Errorf("error while creating product tag: %w", err)
	}
	return nil
}

func (r *productTagRepository) Delete(productID string, tagID int) error {
	row, err := r.db.Exec("DELETE FROM product_tag WHERE product_id = $1 AND tag_id = $2", productID, tagID)
	if rowAffected, _ := row.RowsAffected(); rowAffected == 0 {
		return fmt.Errorf("error while deleting product tag: %w", err)
	} else if err != nil {
		return fmt.Errorf("error while deleting product tag: %w", err)
	}
	return nil
}
