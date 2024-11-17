package domain

import "time"

type Favorite struct {
	USER       *User     `json:"user,omitempty" db:"-"`
	USER_ID    string    `json:"userId" db:"user_id"`
	PRODUCT    *Product  `json:"product,omitempty" db:"-"`
	PRODUCT_ID string    `json:"productId" db:"product_id"`
	CREATED_AT time.Time `json:"createdAt" db:"created_at"`
	UPDATED_AT time.Time `json:"updatedAt" db:"updated_at"`
}

type FavoriteRepository interface {
	GetAll() (*[]Favorite, error)
	GetByID(id string) (*Favorite, error)
	GetByUserID(id string) (*[]Favorite, error)
	GetByProductID(id string) (*[]Favorite, error)
	Create(favorite *Favorite) error
	Delete(user_id string, product_id string) error
}

type FavoriteUsecase interface {
	GetAll() (*[]Favorite, error)
	GetByID(id string) (*Favorite, error)
	GetByUserID(id string) (*[]Favorite, error)
	Create(favorite *Favorite) error
	Delete(user_id string, product_id string) error
}
