package domain

import "github.com/gofiber/fiber/v2"

type AuthUsecase interface {
	SignInWithGoogle(c *fiber.Ctx) (*fiber.Cookie, error)
}
