package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type orderController struct {
	validator      domain.ValidatorUsecase
	orderUsecase   domain.OrderUsecase
	voucherUsecase domain.VoucherUsecase
}

func NewOrderController(validator domain.ValidatorUsecase, orderUsecase domain.OrderUsecase, voucherUsecase domain.VoucherUsecase) *orderController {
	return &orderController{
		validator:      validator,
		orderUsecase:   orderUsecase,
		voucherUsecase: voucherUsecase,
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

// Swagger for create order api
// @Summary Create order
// @Description Create order
// @Tags Order
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body payload.CreateOrderPayload true "Create Order Payload"
// @Success 200 {object} domain.Response
// @Router /api/order [post]
func (c *orderController) CreateOrder(ctx *fiber.Ctx) error {
	var payload payload.CreateOrderPayload
	userID := ctx.Locals(constant.USER_ID).(string)
	if err := c.validator.ValidateBody(ctx, &payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INVALID_BODY,
			SUCCESS: false,
		})
	}

	orderProdcts := make([]domain.OrderProduct, 0)
	totalPrice := float64(0)
	for _, product := range payload.PRODUCTS {
		// IF FOUND SEND ERROR
		productDetail, err := c.orderUsecase.GetProductDetail(product.PRODUCT_SIZE_ID, product.QUANTITY)
		if err != nil || productDetail == nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
				MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
				SUCCESS: false,
			})
		}
		if productDetail.QUANTITY < product.QUANTITY {
			return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
				MESSAGE: constant.MESSAGE_BAD_REQUEST,
				SUCCESS: false,
			})
		}
		totalPrice += productDetail.PRICE * float64(product.QUANTITY)
		orderProdcts = append(orderProdcts, *productDetail)
	}

	// Check Voucher
	var toDiscount float64
	if payload.VOUCHER_ID != "" {
		voucher, err := c.voucherUsecase.GetByID(payload.VOUCHER_ID)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
				MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
				SUCCESS: false,
			})
		}

		if voucher != nil {
			toDiscount = totalPrice * voucher.PERCENTAGE
			// max is voucher max discount
			if toDiscount > voucher.MAX_DISCOUNT {
				toDiscount = voucher.MAX_DISCOUNT
			}
		}
	}

	address := payload.ADDRESS_ID //FUTURE CHANGE TO STRING ADDRESS
	err := c.orderUsecase.CreateOrder(&orderProdcts, address, &payload.VOUCHER_ID, payload.SHIPPING_PRICE, totalPrice, toDiscount, userID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
	})
}

// Swagger for get order detail api that only user can access
// @Summary Get self order detail
// @Description Get self order detail
// @Tags Order
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Order ID"
// @Success 200 {object} domain.Response
// @Error 400 {object} domain.Response
// @Error 404 {object} domain.Response
// @Router /api/order/self/{id} [get]
func (c *orderController) GetSelfOrderDetail(ctx *fiber.Ctx) error {
	userID := ctx.Locals(constant.USER_ID).(string)
	order, err := c.orderUsecase.GetSelfOrderDetail(ctx.Params("id"), userID)
	if order == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_NOT_FOUND,
			SUCCESS: false,
		})
	}
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	return ctx.JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    order,
	})
}
