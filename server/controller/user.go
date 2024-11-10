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

// @Summary Get user profile
// @Description Get user profile
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} domain.Response
// @Router /api/user [get]
func (u *userController) Update(ctx *fiber.Ctx) error {
	var body payload.UpdateUserDTO
	if err := u.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	userID := ctx.Locals(constant.USER_ID).(string)
	if err := u.userUsecase.Update(userID, &domain.User{
		FIRST_NAME:  body.FIRST_NAME,
		LAST_NAME:   body.LAST_NAME,
		PHONE:       body.PHONE,
		PROFILE_URL: body.PROFILE_URL,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}
