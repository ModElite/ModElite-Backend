package domain

import "time"

type ProductStatus string

const (
	ProductActive   ProductStatus = "ACTIVE"
	ProductInactive ProductStatus = "INACTIVE"
	ProductDelete   ProductStatus = "DELETE"
)

type Product struct {
	ID          string    `json:"id" db:"id"`
	SELLER_ID   string    `json:"sellerId" db:"seller_id"`
	NAME        string    `json:"name" db:"name"`
	DESCRIPTION string    `json:"description" db:"description"`
	PRICE       float64   `json:"price" db:"price"`
	STATUS      string    `json:"status" db:"status"`
	CREATED_AT  time.Time `json:"createdAt" db:"created_at"`
	UPDATED_AT  time.Time `json:"updatedAt" db:"updated_at"`
}

type ProductRepository interface {
	GetAll() (*[]Product, error)
	GetByID(id string) (*Product, error)
	GetBySellerID(SellerID string) (*[]Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id string) error
}

type ProductUsecase interface {
	GetAll() (*[]Product, error)
	GetByID(id string) (*Product, error)
	GetBySellerID(SellerID string) (*[]Product, error)
}
