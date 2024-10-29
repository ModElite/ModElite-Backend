package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
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

// @Summary Get all sellers
// @Tags Seller
// @Security Bearer
// @Produce json
// @Success 200 {object} domain.Response
// @Router /api/seller [get]
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

// @Summary Get all sellers by owner
// @Tags Seller
// @Security Bearer
// @Produce json
// @Success 200 {object} domain.Response
// @Router /api/seller/owner [get]
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

// @Summary Get seller by id
// @Tags Seller
// @Security Bearer
// @Produce json
// @Param id path string true "Seller ID"
// @Success 200 {object} domain.Response
// @Router /api/seller/{id} [get]
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

// @Summary Create seller
// @Tags Seller
// @Security Bearer
// @Accept json
// @Produce json
// @Param body body payload.CreateSellerDTO true "Create Seller"
// @Success 201 {object} domain.Response
// @Router /api/seller [post]
func (c *sellerController) Create(ctx *fiber.Ctx) error {
	var body payload.CreateSellerDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	userId := ctx.Locals(constant.USER_ID).(string)
	if err := c.sellerUsecase.Create(&domain.Seller{
		NAME:        body.NAME,
		DESCRIPTION: body.DESCRIPTION,
		LOGO_URL:    body.LOGO_URL,
		LOCATION:    body.LOCATION,
		OWNER_ID:    userId,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}

// @Summary Update seller
// @Tags Seller
// @Security Bearer
// @Accept json
// @Produce json
// @Param body body payload.UpdateSellerDTO true "Update Seller"
// @Success 200 {object} domain.Response
// @Router /api/seller [patch]
func (c *sellerController) Update(ctx *fiber.Ctx) error {
	var body payload.UpdateSellerDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	userId := ctx.Locals(constant.USER_ID).(string)
	err := c.sellerUsecase.Update(body.ID, &domain.Seller{
		NAME:        body.NAME,
		DESCRIPTION: body.DESCRIPTION,
		LOGO_URL:    body.LOGO_URL,
		LOCATION:    body.LOCATION,
	}, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}
