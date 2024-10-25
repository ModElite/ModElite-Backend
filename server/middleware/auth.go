package middleware

import (
	"time"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware(sessionUsecase domain.SessionUsecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ssid := ctx.Cookies(constant.SESSION_COOKIE_NAME)
		if ssid == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: constant.MESSAGE_UNAUTHORIZED,
			})
		}
		session, err := sessionUsecase.GetByID(ssid)
		if session == nil || err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: constant.MESSAGE_UNAUTHORIZED,
			})
		} else if !time.Now().Before(session.EXPIRED_AT) {
			ctx.Cookie(&fiber.Cookie{Name: constant.SESSION_COOKIE_NAME, Expires: time.Unix(0, 0)})
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: constant.MESSAGE_UNAUTHORIZED,
			})
		}

		ctx.Locals(constant.USER_ID, session.USER_ID)

		return ctx.Next()
	}
}
