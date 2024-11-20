package domain

import "time"

type OrderProduct struct {
	ID              string    `json:"id" db:"id"`
	ORDER_ID        string    `json:"orderId" db:"order_id"`
	PRODUCT_SIZE_ID string    `json:"productSizeId" db:"product_size_id"`
	QUANTITY        int       `json:"quantity" db:"quantity"`
	PRICE           float64   `json:"price" db:"price"`
	CREATED_AT      time.Time `json:"createdAt" db:"created_at"`
	UPDATED_AT      time.Time `json:"updatedAt" db:"updated_at"`
	SELLER_ID       string    `json:"sellerId,omitempty" db:"seller_id"`
}

type OrderProductRepository interface {
}

type OrderProductUsecase interface {
}
