package domain

import "time"

type Address struct {
	ID           int       `json:"id" db:"id"`
	USER_ID      string    `json:"userId" db:"user_id"`
	FIRST_NAME   string    `json:"firstName" db:"first_name"`
	LAST_NAME    string    `json:"lastName" db:"last_name"`
	EMAIL        string    `json:"email" db:"email"`
	PHONE        string    `json:"phone" db:"phone"`
	LABEL        string    `json:"label" db:"label"`
	DEFAULT      bool      `json:"default" db:"default"`
	ADDRESS      string    `json:"address" db:"address"`
	SUB_DISTRICT string    `json:"subDistrict" db:"sub_district"`
	DISTRICT     string    `json:"district" db:"district"`
	PROVINCE     string    `json:"province" db:"province"`
	ZIP_CODE     string    `json:"zipCode" db:"zip_code"`
	CREATED_AT   time.Time `json:"createdAt" db:"created_at"`
	UPDATED_AT   time.Time `json:"updatedAt" db:"updated_at"`
}

type AddressRepository interface {
	GetAll() (*[]Address, error)
	GetById(id int) (*Address, error)
	GetByUserId(userId string) (*[]Address, error)
	Create(address *Address) (int, error)
	Update(address *Address) error
	UpdateDefaultByUserId(userId string, id int) error
	Delete(id int) error
}

type AddressUsecase interface {
	CheckPermissionCanModifyAddress(userID string, addressID int) (bool, error)
	GetAll() (*[]Address, error)
	GetAddressByID(addressID int) (*Address, error)
	GetAddressByUserID(userID string) (*[]Address, error)
	Create(address *Address) error
	Update(address *Address) error
	Delete(addressID int) error
}
