package usecase

import (
	"database/sql"
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
)

type cartUsecase struct {
	cartRepository domain.CartRepository
}

func NewCartUsecase(cartRepository domain.CartRepository) domain.CartUsecase {
	return &cartUsecase{
		cartRepository: cartRepository,
	}
}

func (c *cartUsecase) GetAll() (*[]domain.Cart, error) {
	return c.cartRepository.GetAll()
}

func (c *cartUsecase) GetCartByUserId(userId string) (*[]domain.Cart, error) {
	return c.cartRepository.GetCartByUserId(userId)
}

func (c *cartUsecase) EditCart(cart domain.EditCart, userId string) error {
	current_quantity, err := c.cartRepository.GetProductQuantiry(cart.PRODUCT_SIZE_ID)
	if err != nil {
		return err
	}
	if current_quantity < cart.QUANTITY {
		return fmt.Errorf("product quantity in stock is not enough")
	}
	_, err = c.cartRepository.GetCartByUserIdProductSizeID(userId, cart.PRODUCT_SIZE_ID)
	cartItem := domain.Cart{
		USER_ID:         userId,
		PRODUCT_SIZE_ID: cart.PRODUCT_SIZE_ID,
		QUANTITY:        cart.QUANTITY,
	}
	// Product Quantity in Stock to check
	if err == sql.ErrNoRows {
		err := c.cartRepository.AddItemCart(cartItem)
		if err != nil {
			return err
		}
		return nil
	}

	if err != nil {
		return err
	}
	// Update Quantity
	if (cart.QUANTITY) <= 0 {
		err := c.cartRepository.DeleteItemCart(userId, cart.PRODUCT_SIZE_ID)
		if err != nil {
			return err
		}
	} else {
		err := c.cartRepository.UpdateItemCart(userId, cart.PRODUCT_SIZE_ID, cart.QUANTITY)
		if err != nil {
			return err
		}
	}
	return nil

}
