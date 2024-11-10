package controller

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type orderController struct {
	validator    domain.ValidatorUsecase
	orderUsecase domain.OrderUsecase
}

func NewOrderController(validator domain.ValidatorUsecase, orderUsecase domain.OrderUsecase) *orderController {
	return &orderController{
		validator:    validator,
		orderUsecase: orderUsecase,
	}
}

// @Summary Get all order
// @Description Get all order
// @Tags Order
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} domain.Response
// @Router /api/order [get]
func (c *orderController) GetAll(ctx *fiber.Ctx) error {
	orders, err := c.orderUsecase.GetAll()
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    orders,
	})
}

// Swagger for get order api that only user can access
// @Summary Get self order
// @Description Get self order
// @Tags Order
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} domain.Response
// @Router /api/order/self [get]
func (c *orderController) GetSelfOrder(ctx *fiber.Ctx) error {
	userID := ctx.Locals(constant.USER_ID).(string)
	orders, err := c.orderUsecase.GetSelfOrder(userID)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    orders,
	})
}
