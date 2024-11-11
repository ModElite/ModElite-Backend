package domain

type ProductTag struct {
	PRODUCT_ID string `json:"productId" db:"product_id"`
	TAG_ID     int    `json:"tagId" db:"tag_id"`
	TAG        *[]Tag `json:"tag" db:"-"`
}

type ProductTagRepository interface {
	GetAll() (*[]ProductTag, error)
	GetByProductID(productID string) (*[]ProductTag, error)
	GetByTagID(tagID int) (*[]ProductTag, error)
	Create(productID string, tagID int) error
	Delete(productID string, tagID int) error
}
