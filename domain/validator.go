package domain

import "github.com/gofiber/fiber/v2"

type ValidatorUsecase interface {
	ValidateBody(ctx *fiber.Ctx, Schema interface{}) error
}
