package payload

type CreateOrderPayload struct {
	VOUCHER_ID     string                `json:"voucherId"`
	ADDRESS_ID     int                   `json:"addressId" validate:"required"`
	SHIPPING_PRICE float64               `json:"shippingPrice" validate:"required"`
	PRODUCTS       []OrderProductPayload `json:"products" validate:"required"`
}

type OrderProductPayload struct {
	PRODUCT_SIZE_ID string `json:"productSizeId" validate:"required"`
	QUANTITY        int    `json:"quantity" validate:"required"`
}
