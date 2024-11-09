package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type tagsController struct {
	validator   domain.ValidatorUsecase
	tagsUseCase domain.TagsUsecase
}

func NewTagsController(
	validator domain.ValidatorUsecase,
	tagsUseCase domain.TagsUsecase,
) *tagsController {
	return &tagsController{
		validator:   validator,
		tagsUseCase: tagsUseCase,
	}
}

// GetTags godoc
// @Summary Get all tags
// @Description Get all tags
// @Tags Tags
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Response
// @Router /api/tags [get]
func (t *tagsController) GetTags(ctx *fiber.Ctx) error {
	tags, err := t.tagsUseCase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    tags,
	})
}

// CreateTag godoc
// @Summary Create a tag
// @Description Create a tag
// @Tags Tags
// @Accept  json
// @Produce  json
// @Param tag body payload.CreateTagDTO true "Tag"
// @Success 200 {object} domain.Response
// @Router /api/tags [post]
func (t *tagsController) CreateTag(ctx *fiber.Ctx) error {
	var createTagDTO payload.CreateTagDTO
	if err := t.validator.ValidateBody(ctx, &createTagDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INVALID_BODY,
			SUCCESS: false,
		})
	}

	if err := t.tagsUseCase.Create(&domain.Tag{
		LABEL: createTagDTO.LABEL,
		SHOW:  *createTagDTO.SHOW,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
	})
}

// UpdateTag godoc
// @Summary Update a tag
// @Description Update a tag
// @Tags Tags
// @Accept  json
// @Produce  json
// @Param tag body payload.UpdateTagDTO true "Tag"
// @Success 200 {object} domain.Response
// @Router /api/tags [patch]
func (t *tagsController) UpdateTag(ctx *fiber.Ctx) error {
	var updateTagDTO payload.UpdateTagDTO
	if err := t.validator.ValidateBody(ctx, &updateTagDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INVALID_BODY,
			SUCCESS: false,
		})
	}

	if err := t.tagsUseCase.Update(&domain.Tag{
		ID:    updateTagDTO.ID,
		LABEL: updateTagDTO.LABEL,
		SHOW:  *updateTagDTO.SHOW,
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
