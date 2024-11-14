package payload

type AddCartDTO struct {
	PRODUCTSIZEID string `json:"productSizeId" validate:"required"`
	QUANTITY      int    `json:"quantity"`
}
