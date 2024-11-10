package domain

import "time"

type OrderStatus string

const (
	PAYMENTSUCCESS OrderStatus = "PAYMENT_SUCCESS"
	REFUND         OrderStatus = "REFUND"
	END            OrderStatus = "END"
	CANCEL         OrderStatus = "CANCEL"
	PENDING        OrderStatus = "PENDING"
)

type Order struct {
	ID             string      `json:"id" db:"id"`
	USER_ID        string      `json:"user_id" db:"user_id"`
	STATUS         OrderStatus `json:"status" db:"status"`
	TOTAL_PRICE    float64     `json:"total_price" db:"total_price"`
	PRODUCT_PRICE  float64     `json:"product_price" db:"product_price"`
	SHIPPING_PRICE float64     `json:"shipping_price" db:"shipping_price"`
	DISCOUNT       float64     `json:"discount" db:"discount"`
	VOUCHER_CODE   string      `json:"voucher_code" db:"voucher_code"`
	ADDRESS        string      `json:"address" db:"address"`
	CREATED_AT     time.Time   `json:"created_at" db:"created_at"`
	UPDATED_AT     time.Time   `json:"updated_at" db:"updated_at"`
}

type OrderRepository interface {
	GetAll() (*[]Order, error)
	GetSelfOrder(userID string) (*[]Order, error)
}

type OrderUsecase interface {
	GetAll() (*[]Order, error)
	GetSelfOrder(userID string) (*[]Order, error)
}
