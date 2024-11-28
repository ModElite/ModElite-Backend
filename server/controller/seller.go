package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type sellerController struct {
	validator                domain.ValidatorUsecase
	sellerUsecase            domain.SellerUsecase
	sellerTransactionUsecase domain.SellerTransactionUsecase
}

func NewSellerController(validator domain.ValidatorUsecase, sellerUsecase domain.SellerUsecase, sellerTransactionUsecase domain.SellerTransactionUsecase) *sellerController {
	return &sellerController{
		validator:                validator,
		sellerUsecase:            sellerUsecase,
		sellerTransactionUsecase: sellerTransactionUsecase,
	}
}

// @Summary Get all sellers
// @Tags Seller
// @Produce json
// @Success 200 {object} domain.Response{data=[]domain.Seller}
// @Router /api/seller [get]
func (c *sellerController) GetAll(ctx *fiber.Ctx) error {
	sellers, err := c.sellerUsecase.GetAll()
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
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} domain.Response{data=[]domain.Seller}
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

// @Summary Check seller is owner by seller id
// @Tags Seller
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Seller ID"
// @Success 200 {object} domain.Response{data=domain.Seller}
// @Failure 400 {object} domain.Response
// @Failure 403 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/seller/permission/{id} [get]
func (c *sellerController) GetIsOwner(ctx *fiber.Ctx) error {
	userId := ctx.Locals(constant.USER_ID).(string)
	sellerId := ctx.Params("id", "")
	if IsUUID := c.validator.ValidateUUID(sellerId); !IsUUID {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	} else if sellerId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	seller, err := c.sellerUsecase.GetByID(sellerId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	} else if seller == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	} else if seller.OWNER_ID != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_PERMISSION_DENIED,
		})
	}

	sellerTransaction, err := c.sellerTransactionUsecase.GetBySellerId(sellerId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	seller.SELLER_TRANSACTION = sellerTransaction

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    seller,
	})
}

// @Summary Get dashboard by seller id
// @Tags Seller
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Seller ID"
// @Success 200 {object} domain.Response{data=domain.SellerDashboard}
// @Failure 400 {object} domain.Response
// @Failure 403 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/seller/dashboard/{id} [get]
func (c *sellerController) GetDashboard(ctx *fiber.Ctx) error {
	userId := ctx.Locals(constant.USER_ID).(string)
	sellerId := ctx.Params("id", "")
	if IsUUID := c.validator.ValidateUUID(sellerId); !IsUUID {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	seller, err := c.sellerUsecase.GetByID(sellerId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	} else if seller == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	} else if seller.OWNER_ID != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_PERMISSION_DENIED,
		})
	}

	sellerDashboard, err := c.sellerUsecase.GetDashboard(sellerId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    sellerDashboard,
	})

}

// @Summary Get seller by id
// @Tags Seller
// @Produce json
// @Param id path string true "Seller ID"
// @Success 200 {object} domain.Response{data=domain.Seller}
// @Router /api/seller/{id} [get]
func (c *sellerController) GetByID(ctx *fiber.Ctx) error {
	sellerId := ctx.Params("id")

	seller, err := c.sellerUsecase.GetByID(sellerId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: err.Error(),
		})
	} else if seller == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}
	seller.SELLER_TRANSACTION = nil

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    seller,
	})
}

// @Summary Create seller
// @Tags Seller
// @Security ApiKeyAuth
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
		NAME:                  body.NAME,
		DESCRIPTION:           body.DESCRIPTION,
		LOGO_URL:              body.LOGO_URL,
		LOCATION:              body.LOCATION,
		BANK_ACCOUNT_NAME:     body.BANK_ACCOUNT_NAME,
		BANK_ACCOUNT_NUMBER:   body.BANK_ACCOUNT_NUMBER,
		BANK_ACCOUNT_PROVIDER: body.BANK_ACCOUNT_PROVIDER,
		PHONE:                 body.PHONE,
		OWNER_ID:              userId,
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
// @Security ApiKeyAuth
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
		NAME:                  body.NAME,
		DESCRIPTION:           body.DESCRIPTION,
		LOGO_URL:              body.LOGO_URL,
		LOCATION:              body.LOCATION,
		BANK_ACCOUNT_NAME:     body.BANK_ACCOUNT_NAME,
		BANK_ACCOUNT_NUMBER:   body.BANK_ACCOUNT_NUMBER,
		BANK_ACCOUNT_PROVIDER: body.BANK_ACCOUNT_PROVIDER,
		PHONE:                 body.PHONE,
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
