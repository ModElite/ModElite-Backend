package domain

import "time"

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
	PRICE          float64          `json:"price" db:"price"`
	STATUS         string           `json:"status" db:"status"`
	IMAGE_URL      string           `json:"imageUrl" db:"image_url"`
	PRODUCT_OPTION *[]ProductOption `json:"productOption" db:"product_option"`
	CREATED_AT     time.Time        `json:"createdAt" db:"created_at"`
	UPDATED_AT     time.Time        `json:"updatedAt" db:"updated_at"`
}

type ProductRow struct {
	ProductID       string    `db:"id"`
	SellerID        string    `db:"seller_id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	Price           float64   `db:"price"`
	Status          string    `db:"status"`
	ImageURL        string    `db:"image_url"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	OptionID        string    `db:"option_id"`
	OptionLabel     string    `db:"label"`
	OptionImageURL  string    `db:"option_image_url"`
	OptionCreatedAt time.Time `db:"option_created_at"`
	OptionUpdatedAt time.Time `db:"option_updated_at"`
	ProductSizeID   string    `db:"product_size_id"`
	ProductSizeQty  int       `db:"quantity"`
	ProductSizeCA   time.Time `db:"product_size_created_at"`
	ProductSizeUA   time.Time `db:"product_size_updated_at"`
	SizeID          string    `db:"size_id"`
	SizeValue       string    `db:"size"`
	SizeCreatedAt   time.Time `db:"size_created_at"`
	SizeUpdatedAt   time.Time `db:"size_updated_at"`
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
	Delete(id string) error
}

type ProductUsecase interface {
	GetAll() (*[]Product, error)
	GetAllProductWithOptionsAndSizes() (*[]Product, error)
	GetProductWithOptionsAndSizes(productId string) (*Product, error)
	GetProductsBySeller(sellerID string) (*[]Product, error)
	GetByID(id string) (*Product, error)
	GetBySellerID(SellerID string) (*[]Product, error)
}
