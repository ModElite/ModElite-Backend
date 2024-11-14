package domain

import "time"

type Cart struct {
	ID              int          `json:"id" db:"id"`
	USER_ID         string       `json:"userId" db:"user_id"`
	PRODUCT         *ProductData `json:"product,omitempty" db:"-"`
	PRODUCT_SIZE_ID string       `json:"productSizeId" db:"product_size_id"`
	QUANTITY        int          `json:"quantity" db:"quantity"`
	CREATED_AT      time.Time    `json:"createdAt" db:"created_at"`
	UPDATED_AT      time.Time    `json:"updatedAt" db:"updated_at"`
}

type ProductData struct {
	QUANTITY            int    `json:"quantity" db:"quantity"`
	SIZE                string `json:"size" db:"size"`
	PRODUCT_OPTION      string `json:"productOption" db:"product_option"`
	PRODUCT_NAME        string `json:"productName" db:"product_name"`
	PRODUCT_DESCRIPTION string `json:"productDescription" db:"product_description"`
	PRODUCT_PRICE       int    `json:"productPrice" db:"product_price"`
}

type EditCart struct {
	PRODUCT_SIZE_ID string `json:"productSizeId"`
	QUANTITY        int    `json:"quantity"`
}

type CartUsecase interface {
	GetAll() (*[]Cart, error)
	GetCartByUserId(userId string) (*[]Cart, error)
	EditCart(cart EditCart, userId string) error
}

type CartRepository interface {
	GetAll() (*[]Cart, error)
	GetCartByUserId(userId string) (*[]Cart, error)
	GetCartByUserIdProductSizeID(userId string, productSizeId string) (*Cart, error)
	AddItemCart(cart Cart) error
	UpdateItemCart(userId string, productSizeId string, quantity int) error
	DeleteItemCart(userId string, productSizeId string) error
}
