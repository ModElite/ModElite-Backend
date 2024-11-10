package domain

import "time"

type Voucher struct {
	ID              string    `json:"id" db:"id"`
	CODE            string    `json:"code" db:"code"`
	MIN_TOTAL_PRICE float64   `json:"minTotalPrice" db:"min_total_price"`
	MAX_DISCOUNT    float64   `json:"maxDiscount" db:"max_discount"`
	PERCENTAGE      float64   `json:"percentage" db:"percentage"`
	QUOTA           int       `json:"quota" db:"quota"`
	EXPIRED_AT      time.Time `json:"expiredAt" db:"expired_at"`
	UPDATED_AT      time.Time `json:"updatedAt" db:"updated_at"`
	CREATED_AT      time.Time `json:"createdAt" db:"created_at"`
}

type VoucherRepository interface {
	GetByID(id string) (*Voucher, error)
	Search(code string) (*Voucher, error)
	CreateVoucher(voucher *Voucher) error
	CheckDuplicateCode(code string) bool
}

type VoucherUsecase interface {
	GetByID(id string) (*Voucher, error)
	Search(code string) (*Voucher, error)
	CreateVoucher(voucher *Voucher) error
	CheckDuplicateCode(code string) bool
}
