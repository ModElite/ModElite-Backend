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

func (r *productRepository) GetAllProductWithOptionsAndSizes() (*[]domain.Product, error) {
	query := `
		SELECT p.id, p.seller_id, p.name, p.description, p.price, p.status, p.image_url, p.created_at, p.updated_at,
				po.id AS option_id, po.label, po.image_url AS option_image_url, po.created_at AS option_created_at, po.updated_at AS option_updated_at,
				ps.id AS product_size_id, ps.quantity, ps.created_at AS product_size_created_at, ps.updated_at AS product_size_updated_at,
				s.id AS size_id, s.size, s.created_at AS size_created_at, s.updated_at AS size_updated_at
		FROM product p
		LEFT JOIN product_option po ON po.product_id = p.id
		LEFT JOIN product_size ps ON ps.product_option_id = po.id
		LEFT JOIN size s ON s.id = ps.size_id
		ORDER BY p.id, po.id, ps.id
	`

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("error getting product with options and sizes: %w", err)
	}
	defer rows.Close()

	var rowsData []domain.ProductRow
	for rows.Next() {
		var row domain.ProductRow
		if err := rows.StructScan(&row); err != nil {
			return nil, err
		}
		rowsData = append(rowsData, row)
	}

	productsMap := make(map[string]*domain.Product)
	for _, row := range rowsData {
		product, exists := productsMap[row.ProductID]
		if !exists {
			product = &domain.Product{
				ID:             row.ProductID,
				SELLER_ID:      row.SellerID,
				NAME:           row.Name,
				DESCRIPTION:    row.Description,
				PRICE:          row.Price,
				STATUS:         row.Status,
				IMAGE_URL:      row.ImageURL,
				PRODUCT_OPTION: &[]domain.ProductOption{},
				CREATED_AT:     row.CreatedAt,
				UPDATED_AT:     row.UpdatedAt,
			}
			productsMap[row.ProductID] = product
		}
		var option *domain.ProductOption
		for i := range *product.PRODUCT_OPTION {
			if (*product.PRODUCT_OPTION)[i].ID == row.OptionID {
				option = &(*product.PRODUCT_OPTION)[i]
				break
			}
		}
		if option == nil && row.OptionID != "" {
			option = &domain.ProductOption{
				ID:           row.OptionID,
				LABEL:        row.OptionLabel,
				IMAGE_URL:    row.OptionImageURL,
				PRODUCT_SIZE: &[]domain.ProductSize{},
				CREATED_AT:   row.OptionCreatedAt,
				UPDATED_AT:   row.OptionUpdatedAt,
			}
			*product.PRODUCT_OPTION = append(*product.PRODUCT_OPTION, *option)
		}
		if row.ProductSizeID != "" {
			ps := domain.ProductSize{
				ID:                row.ProductSizeID,
				PRODUCT_OPTION_ID: row.OptionID,
				SIZE_ID:           row.SizeID,
				SIZE: &domain.Size{
					ID:         row.SizeID,
					SIZE:       row.SizeValue,
					CREATED_AT: row.SizeCreatedAt,
					UPDATED_AT: row.SizeUpdatedAt,
				},
				QUANTITY:   row.ProductSizeQty,
				CREATED_AT: row.ProductSizeCA,
				UPDATED_AT: row.ProductSizeUA,
			}
			*option.PRODUCT_SIZE = append(*option.PRODUCT_SIZE, ps)
		}
	}
	products := make([]domain.Product, 0, len(productsMap))
	for _, product := range productsMap {
		products = append(products, *product)
	}

	return &products, nil
}

func (r *productRepository) GetProductWithOptionsAndSizes(productId string) (*domain.Product, error) {
	query := `
		SELECT p.id, p.seller_id, p.name, p.description, p.price, p.status, p.image_url, p.created_at, p.updated_at,
		       po.id AS option_id, po.label, po.image_url AS option_image_url, po.created_at AS option_created_at, po.updated_at AS option_updated_at,
		       ps.id AS product_size_id, ps.quantity, ps.created_at AS product_size_created_at, ps.updated_at AS product_size_updated_at,
		       s.id AS size_id, s.size, s.created_at AS size_created_at, s.updated_at AS size_updated_at
		FROM product p
		LEFT JOIN product_option po ON po.product_id = p.id
		LEFT JOIN product_size ps ON ps.product_option_id = po.id
		LEFT JOIN size s ON s.id = ps.size_id
		WHERE p.id = $1
		ORDER BY po.id, ps.id
	`

	rows, err := r.db.Queryx(query, productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product *domain.Product
	productOptionsMap := make(map[string]*domain.ProductOption)

	for rows.Next() {
		var row domain.ProductRow
		if err := rows.StructScan(&row); err != nil {
			return nil, err
		}

		if product == nil {
			product = &domain.Product{
				ID:             row.ProductID,
				SELLER_ID:      row.SellerID,
				NAME:           row.Name,
				DESCRIPTION:    row.Description,
				PRICE:          row.Price,
				STATUS:         row.Status,
				IMAGE_URL:      row.ImageURL,
				PRODUCT_OPTION: &[]domain.ProductOption{},
				CREATED_AT:     row.CreatedAt,
				UPDATED_AT:     row.UpdatedAt,
			}
		}

		if row.OptionID != "" {
			option, exists := productOptionsMap[row.OptionID]
			if !exists {
				option = &domain.ProductOption{
					ID:           row.OptionID,
					LABEL:        row.OptionLabel,
					IMAGE_URL:    row.OptionImageURL,
					PRODUCT_SIZE: &[]domain.ProductSize{},
					CREATED_AT:   row.OptionCreatedAt,
					UPDATED_AT:   row.OptionUpdatedAt,
				}
				productOptionsMap[row.OptionID] = option
				*product.PRODUCT_OPTION = append(*product.PRODUCT_OPTION, *option)
			}

			if row.ProductSizeID != "" {
				productSize := domain.ProductSize{
					ID:                row.ProductSizeID,
					PRODUCT_OPTION_ID: row.OptionID,
					SIZE_ID:           row.SizeID,
					SIZE: &domain.Size{
						ID:         row.SizeID,
						SIZE:       row.SizeValue,
						CREATED_AT: row.SizeCreatedAt,
						UPDATED_AT: row.SizeUpdatedAt,
					},
					QUANTITY:   row.ProductSizeQty,
					CREATED_AT: row.ProductSizeCA,
					UPDATED_AT: row.ProductSizeUA,
				}
				*option.PRODUCT_SIZE = append(*option.PRODUCT_SIZE, productSize)
			}
		}
	}
	return product, nil
}

func (r *productRepository) GetProductsBySeller(sellerID string) (*[]domain.Product, error) {
	query := `
		SELECT p.id, p.seller_id, p.name, p.description, p.price, p.status, p.image_url, p.created_at, p.updated_at,
		       po.id AS option_id, po.label, po.image_url AS option_image_url, po.created_at AS option_created_at, po.updated_at AS option_updated_at,
		       ps.id AS product_size_id, ps.quantity, ps.created_at AS product_size_created_at, ps.updated_at AS product_size_updated_at,
		       s.id AS size_id, s.size, s.created_at AS size_created_at, s.updated_at AS size_updated_at
		FROM product p
		LEFT JOIN product_option po ON po.product_id = p.id
		LEFT JOIN product_size ps ON ps.product_option_id = po.id
		LEFT JOIN size s ON s.id = ps.size_id
		WHERE p.seller_id = $1
		ORDER BY p.id, po.id, ps.id
	`

	rows, err := r.db.Queryx(query, sellerID)
	if err != nil {
		return nil, fmt.Errorf("error getting product with options and sizes: %w", err)
	}
	defer rows.Close()

	var rowsData []domain.ProductRow
	for rows.Next() {
		var row domain.ProductRow
		if err := rows.StructScan(&row); err != nil {
			return nil, err
		}
		rowsData = append(rowsData, row)
	}

	productsMap := make(map[string]*domain.Product)
	for _, row := range rowsData {
		product, exists := productsMap[row.ProductID]
		if !exists {
			product = &domain.Product{
				ID:             row.ProductID,
				SELLER_ID:      row.SellerID,
				NAME:           row.Name,
				DESCRIPTION:    row.Description,
				PRICE:          row.Price,
				STATUS:         row.Status,
				IMAGE_URL:      row.ImageURL,
				PRODUCT_OPTION: &[]domain.ProductOption{},
				CREATED_AT:     row.CreatedAt,
				UPDATED_AT:     row.UpdatedAt,
			}
			productsMap[row.ProductID] = product
		}
		var option *domain.ProductOption
		for i := range *product.PRODUCT_OPTION {
			if (*product.PRODUCT_OPTION)[i].ID == row.OptionID {
				option = &(*product.PRODUCT_OPTION)[i]
				break
			}
		}
		if option == nil && row.OptionID != "" {
			option = &domain.ProductOption{
				ID:           row.OptionID,
				LABEL:        row.OptionLabel,
				IMAGE_URL:    row.OptionImageURL,
				PRODUCT_SIZE: &[]domain.ProductSize{},
				CREATED_AT:   row.OptionCreatedAt,
				UPDATED_AT:   row.OptionUpdatedAt,
			}
			*product.PRODUCT_OPTION = append(*product.PRODUCT_OPTION, *option)
		}
		if row.ProductSizeID != "" {
			ps := domain.ProductSize{
				ID:                row.ProductSizeID,
				PRODUCT_OPTION_ID: row.OptionID,
				SIZE_ID:           row.SizeID,
				SIZE: &domain.Size{
					ID:         row.SizeID,
					SIZE:       row.SizeValue,
					CREATED_AT: row.SizeCreatedAt,
					UPDATED_AT: row.SizeUpdatedAt,
				},
				QUANTITY:   row.ProductSizeQty,
				CREATED_AT: row.ProductSizeCA,
				UPDATED_AT: row.ProductSizeUA,
			}
			*option.PRODUCT_SIZE = append(*option.PRODUCT_SIZE, ps)
		}
	}
	products := make([]domain.Product, 0, len(productsMap))
	for _, product := range productsMap {
		products = append(products, *product)
	}

	return &products, nil
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
