package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type sizeController struct {
	validator   domain.ValidatorUsecase
	sizeUsecase domain.SizeUsecase
}

func NewSizeController(validator domain.ValidatorUsecase, sizeUsecase domain.SizeUsecase) *sizeController {
	return &sizeController{
		validator:   validator,
		sizeUsecase: sizeUsecase,
	}
}

// GetSize godoc
// @Summary Get all size
// @Description Get all size
// @Tags Size
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response{data=domain.Size}
// @Router /api/size [get]
func (s *sizeController) GetSize(ctx *fiber.Ctx) error {
	size, err := s.sizeUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    size,
	})
}

// CreateSize godoc
// @Summary Create size
// @Description Create size
// @Tags Size
// @Accept json
// @Produce json
// @Param size body payload.SizeDTO true "Size body"
// @Success 200 {object} domain.Response
// @Router /api/size [post]
func (s *sizeController) CreateSize(ctx *fiber.Ctx) error {
	var size payload.SizeDTO
	if err := s.validator.ValidateBody(ctx, &size); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INVALID_BODY,
			SUCCESS: false,
		})
	}

	if err := s.sizeUsecase.Create(&domain.Size{
		SIZE: size.SIZE,
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
