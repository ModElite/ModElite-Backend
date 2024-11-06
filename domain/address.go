package domain

import "time"

type Address struct {
	ID             int       `json:"id" db:"id"`
	USER_ID        string    `json:"userId" db:"user_id"`
	FIRST_NAME     string    `json:"firstName" db:"first_name"`
	LAST_NAME      string    `json:"lastName" db:"last_name"`
	COMPANY        string    `json:"company" db:"company"`
	STREET_ADDRESS string    `json:"streetAddress" db:"street_address"`
	STATE          string    `json:"state" db:"state"`
	COUNTRY        string    `json:"country" db:"country"`
	ZIP_CODE       string    `json:"zipCode" db:"zip_code"`
	EMAIL          string    `json:"email" db:"email"`
	PHONE          string    `json:"phone" db:"phone"`
	TYPE           string    `json:"type" db:"type"`
	CREATED_AT     time.Time `json:"createdAt" db:"created_at"`
	UPDATED_AT     time.Time `json:"updatedAt" db:"updated_at"`
}

type AddressRepository interface {
	GetAll() (*[]Address, error)
	GetById(id int) (*Address, error)
	GetByUserId(userId string) (*[]Address, error)
	Create(address *Address) error
	Update(address *Address) error
	Delete(id int) error
}

type AddressUsecase interface {
	GetAll() (*[]Address, error)
	GetAddressByID(addressID int) (*Address, error)
	GetAddressByUserID(userID string) (*[]Address, error)
	Create(address *Address) error
	Update(address *Address) error
	Delete(addressID int) error
}
