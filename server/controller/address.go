package controller

import (
	"strconv"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type addressController struct {
	validator      domain.ValidatorUsecase
	addressUsecase domain.AddressUsecase
	userUsecase    domain.UserUsecase
}

func NewAddressController(
	validator domain.ValidatorUsecase,
	addressUsecase domain.AddressUsecase,
	userUsecase domain.UserUsecase,
) *addressController {
	return &addressController{
		validator:      validator,
		addressUsecase: addressUsecase,
		userUsecase:    userUsecase,
	}
}

// @Summary Get address By User ID
// @Description Get address By User ID
// @Tags Address
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/address [get]
func (a *addressController) GetByUserID(ctx *fiber.Ctx) error {
	address, err := a.addressUsecase.GetAddressByUserID(ctx.Locals(constant.USER_ID).(string))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    address,
	})
}

// @Summary Create Address
// @Description Create Address
// @Tags Address
// @Accept json
// @Produce json
// @Param address body payload.CreateAddressDTO true "Address"
// @Success 200 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/address [post]
func (a *addressController) Create(ctx *fiber.Ctx) error {
	var address payload.CreateAddressDTO
	if err := a.validator.ValidateBody(ctx, &address); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INVALID_BODY,
			SUCCESS: false,
		})
	}

	if err := a.addressUsecase.Create(&domain.Address{
		USER_ID:        ctx.Locals(constant.USER_ID).(string),
		FIRST_NAME:     address.FIRST_NAME,
		LAST_NAME:      address.LAST_NAME,
		COMPANY:        address.COMPANY,
		STREET_ADDRESS: address.STREET_ADDRESS,
		STATE:          address.STATE,
		COUNTRY:        address.COUNTRY,
		ZIP_CODE:       address.ZIP_CODE,
		EMAIL:          address.EMAIL,
		PHONE:          address.PHONE,
		TYPE:           address.TYPE,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
	})
}

// @Summary Update Address
// @Description Update Address
// @Tags Address
// @Accept json
// @Produce json
// @Param address body payload.UpdateAddressDTO true "Address"
// @Success 200 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/address [put]
func (a *addressController) Update(ctx *fiber.Ctx) error {
	var address payload.UpdateAddressDTO
	if err := a.validator.ValidateBody(ctx, &address); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	if err := a.addressUsecase.Update(&domain.Address{
		ID:             address.ID,
		USER_ID:        ctx.Locals(constant.USER_ID).(string),
		FIRST_NAME:     address.FIRST_NAME,
		LAST_NAME:      address.LAST_NAME,
		COMPANY:        address.COMPANY,
		STREET_ADDRESS: address.STREET_ADDRESS,
		STATE:          address.STATE,
		COUNTRY:        address.COUNTRY,
		ZIP_CODE:       address.ZIP_CODE,
		EMAIL:          address.EMAIL,
		PHONE:          address.PHONE,
		TYPE:           address.TYPE,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
	})
}

// @Summary Delete Address
// @Description Delete Address
// @Tags Address
// @Accept json
// @Produce json
// @Param id path int true "Address ID"
// @Success 200 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/address/{id} [delete]
func (a *addressController) Delete(ctx *fiber.Ctx) error {
	addressID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	address, err := a.addressUsecase.GetAddressByID(addressID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	} else if address == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_NOT_FOUND,
			SUCCESS: false,
		})
	}

	if address.USER_ID != ctx.Locals(constant.USER_ID).(string) {
		if isAdmin, err := a.userUsecase.CheckAdmin(ctx.Locals(constant.USER_ID).(string)); err != nil || !isAdmin {
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				MESSAGE: constant.MESSAGE_UNAUTHORIZED,
				SUCCESS: false,
			})
		}
	}

	if err := a.addressUsecase.Delete(addressID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
	})
}
