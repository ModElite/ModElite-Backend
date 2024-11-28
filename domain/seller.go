package domain

import "time"

type Seller struct {
	ID                    string    `json:"id,omitempty" db:"id"`
	NAME                  string    `json:"name" db:"name"`
	DESCRIPTION           string    `json:"description" db:"description"`
	LOGO_URL              string    `json:"logoUrl" db:"logo_url"`
	LOCATION              string    `json:"location" db:"location"`
	BANK_ACCOUNT_NAME     string    `json:"bankAccountName" db:"bank_account_name"`
	BANK_ACCOUNT_NUMBER   string    `json:"bankAccountNumber" db:"bank_account_number"`
	BANK_ACCOUNT_PROVIDER string    `json:"bankAccountProvider" db:"bank_account_provider"`
	PHONE                 string    `json:"phone" db:"phone"`
	OWNER_ID              string    `json:"ownerId" db:"owner_id"`
	IS_VERIFY             bool      `json:"isVerify" db:"is_verify"`
	UPDATED_AT            time.Time `json:"updateAt" db:"updated_at"`
	CREATED_AT            time.Time `json:"createdAt" db:"created_at"`
}

type SellerRepository interface {
	GetAll() (*[]Seller, error)
	GetByID(id string) (*Seller, error)
	GetByOwnerID(ownerID string) (*[]Seller, error)
	Create(seller *Seller) error
	Update(seller *Seller) error
	Delete(id string) error
}

type SellerUsecase interface {
	GetAll() (*[]Seller, error)
	GetByOwner(userId string) (*[]Seller, error)
	GetByID(id string) (*Seller, error)
	Create(seller *Seller) error
	Update(id string, data *Seller, userId string) error
}
