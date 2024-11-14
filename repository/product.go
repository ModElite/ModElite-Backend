package repository

import (
	"database/sql"
	"fmt"

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
		SELECT p.id, p.seller_id, p.name, p.description, p.price, p.status, p.image_url, p.created_at, p.updated_at, p.deleted_at,
				po.id AS option_id, po.label, po.image_url AS option_image_url, po.created_at AS option_created_at, po.updated_at AS option_updated_at, po.deleted_at AS option_deleted_at,
				ps.id AS product_size_id, ps.quantity, ps.created_at AS product_size_created_at, ps.updated_at AS product_size_updated_at, ps.deleted_at AS product_size_deleted_at,
				s.id AS size_id, s.size, s.created_at AS size_created_at, s.updated_at AS size_updated_at
		FROM product p
		LEFT JOIN product_option po ON po.product_id = p.id AND po.deleted_at ISNULL
		LEFT JOIN product_size ps ON ps.product_option_id = po.id AND ps.deleted_at ISNULL
		LEFT JOIN size s ON s.id = ps.size_id
		WHERE p.deleted_at ISNULL
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
				DELETED_AT:     row.DeletedAt,
			}
			productsMap[row.ProductID] = product
		}

		var option *domain.ProductOption
		if row.OptionID.Valid {
			for i := range *product.PRODUCT_OPTION {
				if (*product.PRODUCT_OPTION)[i].ID == row.OptionID.String {
					option = &(*product.PRODUCT_OPTION)[i]
					break
				}
			}
			if option == nil {
				option = &domain.ProductOption{
					ID:           row.OptionID.String,
					LABEL:        row.OptionLabel.String,
					IMAGE_URL:    row.OptionImageURL.String,
					PRODUCT_SIZE: &[]domain.ProductSize{},
					CREATED_AT:   row.OptionCreatedAt.Time,
					UPDATED_AT:   row.OptionUpdatedAt.Time,
					DELETED_AT:   row.OptionDeletedAt,
				}
				*product.PRODUCT_OPTION = append(*product.PRODUCT_OPTION, *option)
			}

			// Add size if available
			if row.ProductSizeID.Valid {
				ps := domain.ProductSize{
					ID:                row.ProductSizeID.String,
					PRODUCT_OPTION_ID: row.OptionID.String,
					SIZE_ID:           row.SizeID.String,
					SIZE: &domain.Size{
						ID:         row.SizeID.String,
						SIZE:       row.SizeValue.String,
						CREATED_AT: row.SizeCreatedAt.Time,
						UPDATED_AT: row.SizeUpdatedAt.Time,
					},
					QUANTITY:   int(row.ProductSizeQty.Int64),
					CREATED_AT: row.ProductSizeCA.Time,
					UPDATED_AT: row.ProductSizeUA.Time,
					DELETED_AT: row.ProductSizeDA,
				}
				*option.PRODUCT_SIZE = append(*option.PRODUCT_SIZE, ps)
			}
		}
	}

	products := make([]domain.Product, 0, len(productsMap))
	for _, product := range productsMap {
		products = append(products, *product)
	}

	return &products, nil
}

func (r *productRepository) GetProductWithOptionsAndSizes(productID string) (*domain.Product, error) {
	query := `
		SELECT p.id, p.seller_id, p.name, p.description, p.price, p.status, p.image_url, p.created_at, p.updated_at, p.deleted_at,
				po.id AS option_id, po.label, po.image_url AS option_image_url, po.created_at AS option_created_at, po.updated_at AS option_updated_at, po.deleted_at AS option_deleted_at,
				ps.id AS product_size_id, ps.quantity, ps.created_at AS product_size_created_at, ps.updated_at AS product_size_updated_at, ps.deleted_at AS product_size_deleted_at,
				s.id AS size_id, s.size, s.created_at AS size_created_at, s.updated_at AS size_updated_at
		FROM product p
		LEFT JOIN product_option po ON po.product_id = p.id AND po.deleted_at ISNULL
		LEFT JOIN product_size ps ON ps.product_option_id = po.id AND ps.deleted_at ISNULL
		LEFT JOIN size s ON s.id = ps.size_id
		WHERE p.id = $1 AND p.deleted_at ISNULL
		ORDER BY po.id, ps.id
	`

	rows, err := r.db.Queryx(query, productID)
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

	var product *domain.Product
	productOptionsMap := make(map[string]*domain.ProductOption)

	for _, row := range rowsData {
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
				DELETED_AT:     row.DeletedAt,
			}
		}

		var option *domain.ProductOption
		if row.OptionID.Valid {
			optionID := row.OptionID.String
			if existingOption, exists := productOptionsMap[optionID]; exists {
				option = existingOption
			} else {
				option = &domain.ProductOption{
					ID:           optionID,
					LABEL:        row.OptionLabel.String,
					IMAGE_URL:    row.OptionImageURL.String,
					PRODUCT_SIZE: &[]domain.ProductSize{},
					CREATED_AT:   row.OptionCreatedAt.Time,
					UPDATED_AT:   row.OptionUpdatedAt.Time,
					DELETED_AT:   row.OptionDeletedAt,
				}
				productOptionsMap[optionID] = option
				*product.PRODUCT_OPTION = append(*product.PRODUCT_OPTION, *option)
			}

			if row.ProductSizeID.Valid {
				ps := domain.ProductSize{
					ID:                row.ProductSizeID.String,
					PRODUCT_OPTION_ID: optionID,
					SIZE_ID:           row.SizeID.String,
					SIZE: &domain.Size{
						ID:         row.SizeID.String,
						SIZE:       row.SizeValue.String,
						CREATED_AT: row.SizeCreatedAt.Time,
						UPDATED_AT: row.SizeUpdatedAt.Time,
					},
					QUANTITY:   int(row.ProductSizeQty.Int64),
					CREATED_AT: row.ProductSizeCA.Time,
					UPDATED_AT: row.ProductSizeUA.Time,
					DELETED_AT: row.ProductSizeDA,
				}
				*option.PRODUCT_SIZE = append(*option.PRODUCT_SIZE, ps)
			}
		}
	}

	if product == nil {
		return nil, nil
	}

	return product, nil
}

func (r *productRepository) GetProductsBySeller(sellerID string) (*[]domain.Product, error) {
	query := `
		SELECT p.id, p.seller_id, p.name, p.description, p.price, p.status, p.image_url, p.created_at, p.updated_at, p.deleted_at,
				po.id AS option_id, po.label, po.image_url AS option_image_url, po.created_at AS option_created_at, po.updated_at AS option_updated_at, po.deleted_at AS option_deleted_at,
				ps.id AS product_size_id, ps.quantity, ps.created_at AS product_size_created_at, ps.updated_at AS product_size_updated_at, ps.deleted_at AS product_size_deleted_at,
				s.id AS size_id, s.size, s.created_at AS size_created_at, s.updated_at AS size_updated_at
		FROM product p
		LEFT JOIN product_option po ON po.product_id = p.id AND po.deleted_at ISNULL
		LEFT JOIN product_size ps ON ps.product_option_id = po.id AND ps.deleted_at ISNULL
		LEFT JOIN size s ON s.id = ps.size_id
		WHERE p.seller_id = $1 AND p.deleted_at ISNULL
		ORDER BY p.id, po.id, ps.id
	`

	rows, err := r.db.Queryx(query, sellerID)
	if err != nil {
		return nil, fmt.Errorf("error getting products by seller: %w", err)
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
				PRODUCT_OPTION: &[]domain.ProductOption{}, // Initialize as empty slice
				CREATED_AT:     row.CreatedAt,
				UPDATED_AT:     row.UpdatedAt,
				DELETED_AT:     row.DeletedAt,
			}
			productsMap[row.ProductID] = product
		}

		// Handle product option if available
		var option *domain.ProductOption
		if row.OptionID.Valid {
			optionID := row.OptionID.String
			// Check if this option already exists for this product
			for i := range *product.PRODUCT_OPTION {
				if (*product.PRODUCT_OPTION)[i].ID == optionID {
					option = &(*product.PRODUCT_OPTION)[i]
					break
				}
			}
			if option == nil {
				// Create and add new option if it doesn't exist
				option = &domain.ProductOption{
					ID:           optionID,
					LABEL:        row.OptionLabel.String,
					IMAGE_URL:    row.OptionImageURL.String,
					PRODUCT_SIZE: &[]domain.ProductSize{},
					CREATED_AT:   row.OptionCreatedAt.Time,
					UPDATED_AT:   row.OptionUpdatedAt.Time,
					DELETED_AT:   row.OptionDeletedAt,
				}
				*product.PRODUCT_OPTION = append(*product.PRODUCT_OPTION, *option)
			}

			// Add product size if available
			if row.ProductSizeID.Valid {
				ps := domain.ProductSize{
					ID:                row.ProductSizeID.String,
					PRODUCT_OPTION_ID: optionID,
					SIZE_ID:           row.SizeID.String,
					SIZE: &domain.Size{
						ID:         row.SizeID.String,
						SIZE:       row.SizeValue.String,
						CREATED_AT: row.SizeCreatedAt.Time,
						UPDATED_AT: row.SizeUpdatedAt.Time,
					},
					QUANTITY:   int(row.ProductSizeQty.Int64),
					CREATED_AT: row.ProductSizeCA.Time,
					UPDATED_AT: row.ProductSizeUA.Time,
					DELETED_AT: row.ProductSizeDA,
				}
				*option.PRODUCT_SIZE = append(*option.PRODUCT_SIZE, ps)
			}
		}
	}

	// Convert map to slice
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
	_, err := r.db.NamedExec("INSERT INTO product (id, seller_id, name, description, price, image_url, status) VALUES (:id, :seller_id, :name, :description, :price, :image_url, :status)", product)
	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}
	return nil
}

func (r *productRepository) Update(product *domain.Product) error {
	_, err := r.db.NamedExec("UPDATE product SET name = :name, description = :description, price = :price, image_url = :image_url, status = :status, updated_at = NOW() WHERE id = :id", product)
	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}
	return nil
}

func (r *productRepository) SoftDelete(id string) error {
	_, err := r.db.Exec("UPDATE product SET deleted_at = NOW() WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error soft deleting product: %w", err)
	}
	return nil
}

func (r *productRepository) SoftDeleteWithOptionsAndSizes(productID string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	// Soft delete product
	_, err = tx.Exec("UPDATE product SET deleted_at = NOW() WHERE id = $1", productID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("error rolling back transaction: %w", rollbackErr)
		}
		return fmt.Errorf("error soft deleting product: %w", err)
	}

	// Soft delete product options
	_, err = tx.Exec("UPDATE product_option SET deleted_at = NOW() WHERE product_id = $1", productID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("error rolling back transaction: %w", rollbackErr)
		}
		return fmt.Errorf("error soft deleting product options: %w", err)
	}

	// Soft delete product sizes
	_, err = tx.Exec(`
		UPDATE 
			product_size ps
		SET
			deleted_at = NOW()
		FROM
			product_option po
		WHERE
			ps.product_option_id = po.id AND po.product_id = $1
		`, productID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("error rolling back transaction: %w", rollbackErr)
		}
		return fmt.Errorf("error soft deleting product sizes: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (r *productRepository) GetProductPriceQuantity(id string) (*domain.ProductPriceQuantity, error) {
	var product domain.ProductPriceQuantity
	err := r.db.Get(&product, `
		SELECT
			product_size.quantity AS quantity,
			product.price AS price,
			product.seller_id AS seller_id
		FROM
			product_size
			INNER JOIN product_option ON product_size.product_option_id = product_option."id"
			INNER JOIN product ON product_option.product_id = product."id" 
		WHERE
			product_size."id" = $1
		LIMIT 1
		`, id)
	if err != nil {
		return nil, fmt.Errorf("error getting product price and quantity: %w", err)
	}
	return &product, nil
}
