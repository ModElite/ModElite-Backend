package controller

import (
	"strconv"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type tagController struct {
	validator  domain.ValidatorUsecase
	tagUseCase domain.TagUsecase
}

func NewTagController(
	validator domain.ValidatorUsecase,
	tagUseCase domain.TagUsecase,
) *tagController {
	return &tagController{
		validator:  validator,
		tagUseCase: tagUseCase,
	}
}

// GetAllTagGroup godoc
// @Summary Get all tag group
// @Description Get all tag group
// @Tags Tag
// @Accept json
// @Produce json
// @Param withTags query bool false "withTag"
// @Success 200 {object} domain.Response
// @Router /api/tag/group [get]
func (c *tagController) GetAllTagGroup(ctx *fiber.Ctx) error {
	withTags, err := strconv.ParseBool(ctx.Query("withTags", "false"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	var tagGroups *[]domain.TagGroup
	if withTags {
		tagGroups, err = c.tagUseCase.GetAllTagGroupWithTags()
	} else {
		tagGroups, err = c.tagUseCase.GetAllTagGroup()
	}
	if err != nil {
		println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    tagGroups,
	})
}

// CreateTagGroup godoc
// @Summary Create tag group
// @Description Create tag group
// @Tags Tag
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body payload.TagGroupDTO true "Create Tag Group"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/tag/group [post]
func (c *tagController) CreateTagGroup(ctx *fiber.Ctx) error {
	var payload payload.TagGroupDTO
	if err := c.validator.ValidateBody(ctx, &payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	tagGroup := &domain.TagGroup{
		LABEL: payload.LABEL,
		SHOW:  *payload.SHOW,
		TAG:   nil,
	}

	if len(payload.TAG) > 0 {
		tag := make([]domain.Tag, 0)
		for _, t := range payload.TAG {
			tag = append(tag, domain.Tag{
				LABEL: t.LABEL,
			})
		}
		tagGroup.TAG = &tag
	}

	if err := c.tagUseCase.CreateTagGroup(tagGroup); err != nil {
		print(err.Error())
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

// UpdateTagGroup godoc
// @Summary Update tag group
// @Description Update tag group
// @Tags Tag
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Tag Group ID"
// @Param body body payload.TagGroupDTO true "Update Tag Group"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/tag/group/{id} [put]
func (c *tagController) UpdateTagGroup(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil || id < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	var tagGroup payload.TagGroupDTO
	if err := c.validator.ValidateBody(ctx, &tagGroup); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	if err := c.tagUseCase.UpdateTagGroup(&domain.TagGroup{
		ID:    int(id),
		LABEL: tagGroup.LABEL,
		SHOW:  *tagGroup.SHOW,
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

// DeleteTagGroup godoc
// @Summary Delete tag group
// @Description Delete tag group
// @Tags Tag
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Tag Group ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/tag/group/{id} [delete]
func (c *tagController) DeleteTagGroup(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil || id < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	if err := c.tagUseCase.DeleteTagGroup(int(id)); err != nil {
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

// GetTag godoc
// @Summary Get tag
// @Description Get tag
// @Tags Tag
// @Accept json
// @Produce json
// @Param tagId query int false "Tag ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 404 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/tag [get]
func (c *tagController) GetTag(ctx *fiber.Ctx) error {
	tagId, err := strconv.ParseInt(ctx.Query("tagId", "0"), 10, 64)
	if err != nil || tagId < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	if tagId == 0 {
		tags, err := c.tagUseCase.GetAllTag()
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
	} else {
		tag, err := c.tagUseCase.GetTagByID(int(tagId))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
				MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
				SUCCESS: false,
			})
		} else if tag == nil {
			return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
				MESSAGE: constant.MESSAGE_NOT_FOUND,
				SUCCESS: false,
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_SUCCESS,
			SUCCESS: true,
			DATA:    tag,
		})
	}
}

// CreateTag godoc
// @Summary Create tag
// @Description Create tag
// @Tags Tag
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body payload.TagDTO true "Create Tag"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/tag [post]
func (c *tagController) CreateTag(ctx *fiber.Ctx) error {
	var tag payload.TagDTO
	if err := c.validator.ValidateBody(ctx, &tag); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	if err := c.tagUseCase.CreateTag(&domain.Tag{
		TAG_GRUOP_ID: tag.TAG_GROUP_ID,
		LABEL:        tag.LABEL,
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

// UpdateTag godoc
// @Summary Update tag
// @Description Update tag
// @Tags Tag
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Tag ID"
// @Param body body payload.TagDTO true "Update Tag"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/tag/{id} [put]
func (c *tagController) UpdateTag(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil || id < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}
	var tag payload.TagDTO
	if err := c.validator.ValidateBody(ctx, &tag); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	if err := c.tagUseCase.UpdateTag(&domain.Tag{
		ID:           int(id),
		TAG_GRUOP_ID: tag.TAG_GROUP_ID,
		LABEL:        tag.LABEL,
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

// DeleteTag godoc
// @Summary Delete tag
// @Description Delete tag
// @Tags Tag
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Tag ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /api/tag/{id} [delete]
func (c *tagController) DeleteTag(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil || id < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}

	if err := c.tagUseCase.DeleteTag(int(id)); err != nil {
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