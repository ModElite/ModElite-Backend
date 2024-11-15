package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	validator   domain.ValidatorUsecase
	userUsecase domain.UserUsecase
}

func NewUserController(
	validator domain.ValidatorUsecase,
	userUsecase domain.UserUsecase,
) *userController {
	return &userController{
		validator:   validator,
		userUsecase: userUsecase,
	}
}

// @Summary Patch user profile
// @Description Patch user profile
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body payload.UpdateInfoUserDTO true "Patch user profile"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/user [patch]
func (u *userController) UpdateInfo(ctx *fiber.Ctx) error {
	var body payload.UpdateInfoUserDTO
	if err := u.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	userID := ctx.Locals(constant.USER_ID).(string)
	if err := u.userUsecase.UpdateInfo(&domain.User{
		ID:         userID,
		FIRST_NAME: body.FIRST_NAME,
		LAST_NAME:  body.LAST_NAME,
		PHONE:      body.PHONE,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	data, err := u.userUsecase.Get(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    data,
	})
}

// @Summary Patch profile
// @Description Patch profile
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body payload.UpdateImageUserDTO true "Patch user profile"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/user/profile [patch]
func (u *userController) UpdateImage(ctx *fiber.Ctx) error {
	var body payload.UpdateImageUserDTO
	if err := u.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	userID := ctx.Locals(constant.USER_ID).(string)
	if err := u.userUsecase.UpdateImage(&domain.User{
		ID:          userID,
		PROFILE_URL: body.PROFILE_URL,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	data, err := u.userUsecase.Get(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    data,
	})
}
