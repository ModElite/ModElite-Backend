package domain

import "time"

type Size struct {
	ID         string    `json:"id" db:"id"`
	SIZE       string    `json:"size" db:"size"`
	CREATED_AT time.Time `json:"createdAt" db:"created_at"`
	UPDATED_AT time.Time `json:"updatedAt" db:"updated_at"`
}

type SizeRepository interface {
	GetAll() (*[]Size, error)
	GetByID(id string) (*Size, error)
	Create(size *Size) error
	Update(size *Size) error
	Delete(id string) error
}

type SizeUsecase interface {
	GetAll() (*[]Size, error)
	GetByID(id string) (*Size, error)
	Create(size *Size) error
	Update(size *Size) error
	Delete(id string) error
}
