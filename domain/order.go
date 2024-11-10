package domain

import "time"

type OrderStatusType string

const (
	ORDER_PENDING         OrderStatusType = "PENDING"
	ORDER_PAYMENT_SUCCESS OrderStatusType = "PAYMENT_SUCCESS"
	ORDER_REFUND          OrderStatusType = "REFUND"
	ORDER_END             OrderStatusType = "END"
	ORDER_CANCEL          OrderStatusType = "CANCEL"
)

type Order struct {
	ID             string          `json:"id" db:"id"`
	ORDER_PRODUCT  *[]OrderProduct `json:"orderProduct,omitempty" db:"-"`
	USER           *User           `json:"user,omitempty" db:"-"`
	USER_ID        string          `json:"userId" db:"user_id"`
	STATUS         OrderStatusType `json:"status" db:"status"`
	TOTAL_PRICE    float64         `json:"totalPrice" db:"total_price"`
	PRODUCT_PRICE  float64         `json:"productPrice" db:"product_price"`
	SHIPPING_PRICE float64         `json:"shippingPrice" db:"shipping_price"`
	DISCOUNT       float64         `json:"discount" db:"discount"`
	VOUCHER_CODE   *string         `json:"voucherCode" db:"voucher_code"`
	ADDRESS        string          `json:"address" db:"address"`
	CREATED_AT     time.Time       `json:"createdAt" db:"created_at"`
	UPDATED_AT     time.Time       `json:"updatedAt" db:"updated_at"`
}

type OrderRepository interface {
	GetAll() (*[]Order, error)
	GetSelfOrder(userID string) (*[]Order, error)
	CreateOrder(order *[]OrderProduct, address string, voucherId *string, shipping_price float64, totalPrice float64, toDiscount float64, userId string) error
}

type OrderUsecase interface {
	GetAll() (*[]Order, error)
	GetSelfOrder(userID string) (*[]Order, error)
	CreateOrder(order *[]OrderProduct, address string, voucherId *string, shipping_price float64, totalPrice float64, toDiscount float64, userId string) error
	GetProductDetail(productSizeID string, quantity int) (*OrderProduct, error)
}
