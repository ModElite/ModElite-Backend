package payload

type CreateProductDTO struct {
	SELLER_ID      string              `json:"sellerId" validate:"required"`
	NAME           string              `json:"name" validate:"required"`
	DESCRIPTION    string              `json:"description" validate:"required"`
	FEATURE        string              `json:"feature" validate:"required"`
	IMAGE_URL      string              `json:"imageUrl" validate:"required"`
	PRICE          float64             `json:"price" validate:"required,min=0"`
	PRODUCT_OPTION *[]ProductOptionDTO `json:"productOption" validate:"required"`
	TAG_ID         *[]int              `json:"tagId" validate:"omitempty"`
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

type UpdateProductDTO struct {
	NAME           string                    `json:"name" validate:"required"`
	DESCRIPTION    string                    `json:"description" validate:"required"`
	FEATURE        string                    `json:"feature" validate:"required"`
	PRICE          float64                   `json:"price" validate:"required,min=0"`
	IMAGE_URL      string                    `json:"imageUrl" validate:"omitempty"`
	PRODUCT_OPTION *[]UpdateProductOptionDTO `json:"productOption" validate:"required"`
	TAG_ID         *[]int                    `json:"tagId" validate:"omitempty"`
}

type UpdateProductOptionDTO struct {
	ID           string                  `json:"id" validate:"omitempty"`
	LABEL        string                  `json:"label" validate:"required"`
	IMAGE_URL    string                  `json:"imageUrl" validate:"omitempty"`
	PRODUCT_SIZE *[]UpdateProductSizeDTO `json:"productSize" validate:"required"`
}

type UpdateProductSizeDTO struct {
	ID       string `json:"id" validate:"omitempty"`
	SIZE_ID  string `json:"sizeId" validate:"omitempty"`
	QUANTITY int    `json:"quantity" validate:"required"`
}

type FilterDTO struct {
	Filter []FilterTagsDTO `json:"filter"`
}
type FilterTagsDTO struct {
	Name  string   `json:"name"`
	Value []string `json:"value"`
}
