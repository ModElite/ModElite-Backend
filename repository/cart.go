package repository

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type cartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) domain.CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (c *cartRepository) GetAll() (*[]domain.Cart, error) {
	var carts []domain.Cart
	err := c.db.Select(&carts, "SELECT * FROM cart")
	if err != nil {
		return nil, err
	}
	return &carts, nil
}

func (c *cartRepository) GetCartByUserId(userId string) (*[]domain.Cart, error) {
	var carts []domain.Cart
	err := c.db.Select(&carts, "SELECT * FROM cart WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	for i := range carts {
		product := domain.ProductData{}
		err := c.db.Get(&product, `
			SELECT
				product_size.quantity as quantity,
				"size"."size" as size,
				product_option.label as product_option,
				product."name" as product_name,
				product.description as product_description,
				product.price as product_price
			FROM
				product_size
				INNER JOIN product_option ON product_size.product_option_id = product_option."id"
				INNER JOIN product ON product_option.product_id = product."id"
				INNER JOIN "size" ON product_size.size_id = "size"."id"
			WHERE 
				product_size."id" = $1
			LIMIT 1;
		`, carts[i].PRODUCT_SIZE_ID)
		carts[i].PRODUCT = &product
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &carts, nil
}

func (c *cartRepository) GetCartByUserIdProductSizeID(userId string, productSizeId string) (*domain.Cart, error) {
	var cart domain.Cart
	err := c.db.Get(&cart, "SELECT * FROM cart WHERE user_id = $1 AND product_size_id = $2", userId, productSizeId)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (c *cartRepository) AddItemCart(cart domain.Cart) error {
	_, err := c.db.Exec("INSERT INTO cart (user_id, product_size_id, quantity) VALUES ($1, $2, $3)", cart.USER_ID, cart.PRODUCT_SIZE_ID, cart.QUANTITY)
	if err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) DeleteItemCart(userId string, productSizeId string) error {
	_, err := c.db.Exec("DELETE FROM cart WHERE user_id = $1 AND product_size_id = $2", userId, productSizeId)
	if err != nil {
		return err
	}
	return nil
}

// UpdateItemCart
func (c *cartRepository) UpdateItemCart(userId string, productSizeId string, quantity int) error {
	_, err := c.db.Exec("UPDATE cart SET quantity = $1 WHERE user_id = $2 AND product_size_id = $3", quantity, userId, productSizeId)
	if err != nil {
		return err
	}
	return nil
}