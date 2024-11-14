package controller

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type cartController struct {
	validator   domain.ValidatorUsecase
	cartUsecase domain.CartUsecase
}

func NewCartController(validator domain.ValidatorUsecase, cartUsecase domain.CartUsecase) *cartController {
	return &cartController{
		validator:   validator,
		cartUsecase: cartUsecase,
	}
}

// @Summary Get all cart
// @Description Get all cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} domain.Response{data=[]domain.Cart}
// @Router /api/cart [get]
func (c *cartController) GetAll(ctx *fiber.Ctx) error {
	carts, err := c.cartUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    carts,
	})
}

// @Summary Get cart by user id
// @Description Get cart by user id
// @Tags Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} domain.Response{data=[]domain.Cart}
// @Router /api/cart/self [get]
func (c *cartController) GetCartByUserId(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)
	carts, err := c.cartUsecase.GetCartByUserId(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    carts,
	})
}

// @Summary Add cart
// @Description Add cart or update cart quantity if product already in cart for quantity is 0 then delete cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body payload.AddCartDTO true "Add cart"
// @Success 200 {object} domain.Response{data=[]domain.Cart}
// @Router /api/cart [post]
func (c *cartController) EditCart(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)
	var addCardData payload.AddCartDTO
	if err := c.validator.ValidateBody(ctx, &addCardData); err != nil {
		return err
	}

	editCart := domain.EditCart{
		PRODUCT_SIZE_ID: addCardData.PRODUCTSIZEID,
		QUANTITY:        addCardData.QUANTITY,
	}

	err := c.cartUsecase.EditCart(editCart, userID)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}
	carts, err := c.cartUsecase.GetCartByUserId(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    carts,
	})
}
