package validator

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type payloadValidator struct {
	validator *validator.Validate
}

func NewPayloadValidator() domain.ValidatorUsecase {
	return &payloadValidator{
		validator: validator.New(),
	}
}

func (v *payloadValidator) ValidateBody(ctx *fiber.Ctx, Schema interface{}) error {
	ctx.BodyParser(Schema)
	if err := v.validator.Struct(Schema); err != nil {
		return err
	}

	return nil
}
