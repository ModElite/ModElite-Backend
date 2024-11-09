package middleware

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/gofiber/fiber/v2"
)

func NewAuthAdminMiddleware(userUsercase domain.UserUsecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID := ctx.Locals(constant.USER_ID).(string)

		if isAdmin, err := userUsercase.CheckAdmin(userID); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
				MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
				SUCCESS: false,
			})
		} else if !isAdmin {
			return ctx.Status(fiber.StatusForbidden).JSON(domain.Response{
				MESSAGE: constant.MESSAGE_PERMISSION_DENIED,
				SUCCESS: false,
			})
		}

		return ctx.Next()
	}
}
