package domain

import "time"

type ProductOption struct {
	ID           string         `json:"id" db:"id"`
	PRODUCT_ID   string         `json:"-" db:"product_id"`
	PRODUCT_SIZE *[]ProductSize `json:"productSize" db:"-"`
	LABEL        string         `json:"label" db:"label"`
	IMAGE_URL    string         `json:"imageUrl" db:"image_url"`
	CREATED_AT   time.Time      `json:"createdAt" db:"created_at"`
	UPDATED_AT   time.Time      `json:"updatedAt" db:"updated_at"`
}

type ProductOptionRepository interface {
	GetAll() (*[]ProductOption, error)
	GetByID(id string) (*ProductOption, error)
	GetByProductID(productID string) (*[]ProductOption, error)
	Create(productOption *ProductOption) error
	Update(productOption *ProductOption) error
}

type ProductOptionUsecase interface {
	GetAll() (*[]ProductOption, error)
	GetByID(id string) (*ProductOption, error)
	GetByProductID(productID string) (*[]ProductOption, error)
	Create(productOption *ProductOption) error
	Update(productOption *ProductOption) error
}
