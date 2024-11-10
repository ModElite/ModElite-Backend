package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type voucherController struct {
	validator      domain.ValidatorUsecase
	voucherUsecase domain.VoucherUsecase
}

func NewVoucherController(validator domain.ValidatorUsecase, voucherUsecase domain.VoucherUsecase) *voucherController {
	return &voucherController{
		validator:      validator,
		voucherUsecase: voucherUsecase,
	}
}

// @Summary Search voucher by code
// @Description Search voucher by code
// @Tags Voucher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param code path string true "Code"
// @Success 200 {object} domain.Response
// @Router /api/voucher/{code} [get]
func (c *voucherController) Search(ctx *fiber.Ctx) error {
	code := ctx.Params("code", "")
	if code == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: "Code is required",
			SUCCESS: false,
		})
	}

	voucher, err := c.voucherUsecase.Search(code)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_NOT_FOUND,
			SUCCESS: false,
		})
	}

	return ctx.JSON(domain.Response{
		MESSAGE: "Success",
		SUCCESS: true,
		DATA:    voucher,
	})
}

// @Summary Create voucher
// @Description Create voucher
// @Tags Voucher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param voucher body payload.CreateVoucherDTO true "Voucher"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Router /api/voucher [post]
func (c *voucherController) CreateVoucher(ctx *fiber.Ctx) error {
	var payload payload.CreateVoucherDTO
	if err := c.validator.ValidateBody(ctx, &payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INVALID_BODY,
			SUCCESS: false,
		})
	}

	CheckDuplicateCode := c.voucherUsecase.CheckDuplicateCode(payload.Code)
	if CheckDuplicateCode {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: "Code already exist",
			SUCCESS: false,
		})
	}

	voucher := &domain.Voucher{
		CODE:            payload.Code,
		MIN_TOTAL_PRICE: payload.MinTotalPrice,
		MAX_DISCOUNT:    payload.MaxDiscount,
		PERCENTAGE:      payload.Percentage,
		QUOTA:           payload.Quota,
		EXPIRED_AT:      payload.ExpiredAt,
	}

	if err := c.voucherUsecase.CreateVoucher(voucher); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.JSON(domain.Response{
		MESSAGE: "Success",
		SUCCESS: true,
	})
}
