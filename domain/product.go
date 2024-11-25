package domain

import (
	"database/sql"
	"time"
)

type ProductStatus string

const (
	ProductActive   ProductStatus = "ACTIVE"
	ProductInactive ProductStatus = "INACTIVE"
	ProductDelete   ProductStatus = "DELETE"
)

type Product struct {
	ID             string           `json:"id" db:"id"`
	SELLER_ID      string           `json:"sellerId" db:"seller_id"`
	NAME           string           `json:"name" db:"name"`
	DESCRIPTION    string           `json:"description" db:"description"`
	FEATURE        string           `json:"feature" db:"feature"`
	PRICE          float64          `json:"price" db:"price"`
	STATUS         string           `json:"status" db:"status"`
	IMAGE_URL      string           `json:"imageUrl" db:"image_url"`
	PRODUCT_OPTION *[]ProductOption `json:"productOption" db:"product_option"`
	TAGS           *[]Tag           `json:"tags" db:"-"`
	SELLER         *Seller          `json:"seller,omitempty" db:"-"`
	SELLER_NAME    *string          `json:"seller_name,omitempty" db:"seller_name"`
	CREATED_AT     time.Time        `json:"createdAt" db:"created_at"`
	UPDATED_AT     time.Time        `json:"updatedAt" db:"updated_at"`
	DELETED_AT     *time.Time       `json:"deletedAt" db:"deleted_at"`
}

type ProductRow struct {
	ProductID       string         `db:"id"`
	SellerID        string         `db:"seller_id"`
	Name            string         `db:"name"`
	Description     string         `db:"description"`
	Feature         string         `db:"feature"`
	Price           float64        `db:"price"`
	Status          string         `db:"status"`
	SellerName      string         `db:"seller_name"`
	ImageURL        string         `db:"image_url"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
	DeletedAt       *time.Time     `db:"deleted_at"`
	OptionID        sql.NullString `db:"option_id"`
	OptionLabel     sql.NullString `db:"label"`
	OptionImageURL  sql.NullString `db:"option_image_url"`
	OptionCreatedAt sql.NullTime   `db:"option_created_at"`
	OptionUpdatedAt sql.NullTime   `db:"option_updated_at"`
	OptionDeletedAt *time.Time     `db:"option_deleted_at"`
	ProductSizeID   sql.NullString `db:"product_size_id"`
	ProductSizeQty  sql.NullInt64  `db:"quantity"`
	ProductSizeCA   sql.NullTime   `db:"product_size_created_at"`
	ProductSizeUA   sql.NullTime   `db:"product_size_updated_at"`
	ProductSizeDA   *time.Time     `db:"product_size_deleted_at"`
	SizeID          sql.NullString `db:"size_id"`
	SizeValue       sql.NullString `db:"size"`
	SizeCreatedAt   sql.NullTime   `db:"size_created_at"`
	SizeUpdatedAt   sql.NullTime   `db:"size_updated_at"`
}

type ProductPriceQuantity struct {
	Price    float64 `json:"price" db:"price"`
	Quantity int     `json:"quantity" db:"quantity"`
	SellerID string  `json:"sellerId" db:"seller_id"`
}

type ProductRepository interface {
	GetAllProductWithOptionsAndSizes() (*[]Product, error)
	GetProductWithOptionsAndSizes(productId string) (*Product, error)
	GetProductsBySeller(sellerID string) (*[]Product, error)
	GetAll() (*[]Product, error)
	GetByID(id string) (*Product, error)
	GetBySellerID(SellerID string) (*[]Product, error)
	Create(product *Product) error
	Update(product *Product) error
	SoftDelete(id string) error
	SoftDeleteWithOptionsAndSizes(productID string) error
	GetProductPriceQuantity(productSizeID string) (*ProductPriceQuantity, error)
}

type ProductUsecase interface {
	CheckPermissionCanModifyProduct(ownerID string, productID string) (bool, error)
	GetAll() (*[]Product, error)
	GetAllProductWithOptionsAndSizes() (*[]Product, error)
	GetProductWithOptionsAndSizes(productId string) (*Product, error)
	GetProductsBySeller(sellerID string) (*[]Product, error)
	GetByID(id string) (*Product, error)
	GetBySellerID(SellerID string) (*[]Product, error)
	Create(product *Product) (id *string, err error)
	Update(newProduct *Product) error
	SoftDeleteWithOptionsAndSizes(id string) error
}
