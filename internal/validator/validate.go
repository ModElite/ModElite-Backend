package validator

import (
	"regexp"

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
	if err := ctx.BodyParser(Schema); err != nil {
		return err
	}
	if err := v.validator.Struct(Schema); err != nil {
		return err
	}

	return nil
}

func (v *payloadValidator) ValidateUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
