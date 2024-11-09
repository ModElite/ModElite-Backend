package domain

import "time"

type Tag struct {
	ID         int       `json:"id" db:"id"`
	LABEL      string    `json:"label" db:"label"`
	SHOW       bool      `json:"show" db:"show"`
	CREATED_AT time.Time `json:"createdAt" db:"created_at"`
}

type TagsRepository interface {
	GetAll() (*[]Tag, error)
	GetByID(id int) (*Tag, error)
	Create(tag *Tag) error
	Update(tag *Tag) error
	Delete(id int) error
}

type TagsUsecase interface {
	GetAll() (*[]Tag, error)
	GetByID(id int) (*Tag, error)
	GetByProductID(productId string) (*[]Tag, error)
	Create(tag *Tag) error
	Update(tag *Tag) error
	Delete(id int) error
}
