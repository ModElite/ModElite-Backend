package domain

import "time"

type ProductSize struct {
	ID                string    `json:"id" db:"id"`
	PRODUCT_OPTION_ID string    `json:"-" db:"product_option_id"`
	SIZE_ID           string    `json:"-" db:"size_id"`
	SIZE              *Size     `json:"size" db:"-"`
	QUANTITY          int       `json:"quantity" db:"quantity"`
	CREATED_AT        time.Time `json:"createdAt" db:"created_at"`
	UPDATED_AT        time.Time `json:"updatedAt" db:"updated_at"`
}

type ProductSizeRepository interface {
	GetAll() (*[]ProductSize, error)
	GetByID(id string) (*ProductSize, error)
	GetByProductOptionID(productOptionID string) (*[]ProductSize, error)
	Create(productSize *ProductSize) error
	Update(productSize *ProductSize) error
}

type ProductSizeUsecase interface {
	GetAll() (*[]ProductSize, error)
	GetByID(id string) (*ProductSize, error)
	GetByProductOptionID(productOptionID string) (*[]ProductSize, error)
	Create(productSize *ProductSize) error
	Update(productSize *ProductSize) error
}
