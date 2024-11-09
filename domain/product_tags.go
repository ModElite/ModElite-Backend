package domain

type ProductTag struct {
	ID         int    `json:"id" db:"id"`
	PRODUCT_ID string `json:"productId" db:"product_id"`
	TAG_ID     int    `json:"tagId" db:"tag_id"`
	TAG        *[]Tag `json:"tag" db:"-"`
	CREATED_AT string `json:"createdAt" db:"created_at"`
	UPDATED_AT string `json:"updatedAt" db:"updated_at"`
}

type ProductTagsRepository interface {
	GetAll() (*[]ProductTag, error)
	GetByProductID(productID string) (*[]ProductTag, error)
	GetByTagID(tagID int) (*[]ProductTag, error)
	Create(productID string, tagID int) error
	Delete(productID string, tagID int) error
}
