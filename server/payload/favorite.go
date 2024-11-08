package payload

type CreateFavoriteDTO struct {
	PRODUCT_ID string `json:"productId" validate:"required"`
}
