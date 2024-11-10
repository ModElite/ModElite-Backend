package payload

import "time"

type CreateVoucherDTO struct {
	Code          string    `json:"code" validate:"required"`
	MinTotalPrice float64   `json:"minTotalPrice" validate:"required"`
	MaxDiscount   float64   `json:"maxDiscount" validate:"required"`
	Percentage    float64   `json:"percentage" validate:"required"`
	Quota         int       `json:"quota" validate:"required"`
	ExpiredAt     time.Time `json:"expiredAt" validate:"required"`
}
