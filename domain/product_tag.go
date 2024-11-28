package domain

type ProductTag struct {
	PRODUCT_ID string `json:"productId" db:"product_id"`
	TAG_ID     int    `json:"tagId" db:"tag_id"`
	TAG        *Tag   `json:"tag" db:"-"`
}

type ProductTagJoinTagRow struct {
	ProductID   string `json:"productId" db:"product_id"`
	TagID       int    `json:"tagId" db:"tag_id"`
	TagLabel    string `json:"label" db:"tag_label"`
	TagImageURL string `json:"imageUrl" db:"tag_image_url"`
}

type ProductTagRepository interface {
	GetAll() (*[]ProductTag, error)
	GetAllJoinTag() (*[]ProductTag, error)
	GetByProductID(productID string) (*[]ProductTag, error)
	GetByTagID(tagID int) (*[]ProductTag, error)
	Create(productID string, tagID int) error
	Delete(productID string, tagID int) error
}
