package domain

import "time"

type OrderProductStatus string

const (
	ORDER_PRODUCT_PENDING             OrderProductStatus = "PENDING"
	ORDER_PRODUCT_DELIVERY_ON_THE_WAY OrderProductStatus = "DELIVERY_ON_THE_WAY"
	ORDER_PRODUCT_DELIVERED                              = "DELIVERED"
)

type OrderProduct struct {
	ID              string             `json:"id" db:"id"`
	ORDER_ID        string             `json:"orderId" db:"order_id"`
	PRODUCT_SIZE_ID string             `json:"productSizeId" db:"product_size_id"`
	STATUS          OrderProductStatus `json:"status" db:"status"`
	QUANTITY        int                `json:"quantity" db:"quantity"`
	PRICE           float64            `json:"price" db:"price"`
	CREATED_AT      time.Time          `json:"createdAt" db:"created_at"`
	UPDATED_AT      time.Time          `json:"updatedAt" db:"updated_at"`
	SELLER_ID       string             `json:"sellerId,omitempty" db:"seller_id"`
}

type OrderProductRepository interface {
}

type OrderProductUsecase interface {
}
