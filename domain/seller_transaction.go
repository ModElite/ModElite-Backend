package domain

import "time"

type SellerTransaction struct {
	ID                        string                    `json:"id" db:"id"`
	SELLER_ID                 string                    `json:"sellerId" db:"seller_id"`
	BANK_ACCOUNT_NAME         string                    `json:"bankAccountName" db:"bank_account_name"`
	BANK_ACCOUNT_NUMBER       string                    `json:"bankAccountNumber" db:"bank_account_number"`
	BANK_ACCOUNT_PROVIDER     string                    `json:"bankAccountProvider" db:"bank_account_provider"`
	BANK_TRANSACTION_ID       string                    `json:"bankTransactionId" db:"bank_transaction_id"`
	BANK_TRANSACTION_AMOUNT   float64                   `json:"bankTransactionAmount" db:"bank_transaction_amount"`
	BANK_TRANSACTION_FEE      float64                   `json:"bankTransactionFee" db:"bank_transaction_fee"`
	BANK_TRANSACTION_DATETIME time.Time                 `json:"bankTransactionDatetime" db:"bank_transaction_datetime"`
	CREATED_AT                time.Time                 `json:"createdAt" db:"created_at"`
	UPDATED_AT                time.Time                 `json:"updatedAt" db:"updated_at"`
	SELLER_TRANSACTION_ORDER  *[]SellerTransactionOrder `json:"sellerTransactionOrder,omitempty" db:"-"`
}

type SellerTransactionRepository interface {
	GetAll() (*[]SellerTransaction, error)
	GetBySellerId(sellerId string) (*[]SellerTransaction, error)
}

type SellerTransactionUsecase interface {
	GetAll() (*[]SellerTransaction, error)
	GetBySellerId(sellerId string) (*[]SellerTransaction, error)
}
