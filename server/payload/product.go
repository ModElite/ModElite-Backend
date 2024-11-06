package payload

type CreateProductDTO struct {
	SELLER_ID      string              `json:"sellerId" validate:"required"`
	NAME           string              `json:"name" validate:"required"`
	DESCRIPTION    string              `json:"description" validate:"required"`
	PRICE          float64             `json:"price" validate:"required"`
	IMAGE_URL      string              `json:"imageUrl" validate:"required"`
	PRODUCT_OPTION *[]ProductOptionDTO `json:"productOption" validate:"required"`
}

type ProductOptionDTO struct {
	LABEL        string            `json:"label" validate:"required"`
	IMAGE_URL    string            `json:"imageUrl" validate:"required"`
	PRODUCT_SIZE *[]ProductSizeDTO `json:"productSize" validate:"required"`
}

type ProductSizeDTO struct {
	SIZE_ID  string `json:"sizeId" validate:"required"`
	QUANTITY int    `json:"quantity" validate:"required"`
}
