package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type sellerController struct {
	validator     domain.ValidatorUsecase
	sellerUsecase domain.SellerUsecase
}

func NewSellerController(validator domain.ValidatorUsecase, sellerUsecase domain.SellerUsecase) *sellerController {
	return &sellerController{
		validator:     validator,
		sellerUsecase: sellerUsecase,
	}
}

func (c *sellerController) GetAll(ctx *fiber.Ctx) error {
	sellers, err := c.sellerUsecase.GetAll(ctx.Locals(constant.USER_ID).(string))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    sellers,
	})
}

func (c *sellerController) GetByOwner(ctx *fiber.Ctx) error {
	userId := ctx.Locals(constant.USER_ID).(string)

	sellers, err := c.sellerUsecase.GetByOwner(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    sellers,
	})
}

func (c *sellerController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userId := ctx.Locals(constant.USER_ID).(string)

	seller, err := c.sellerUsecase.GetByID(id, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    seller,
	})
}
