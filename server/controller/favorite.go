package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type favoriteController struct {
	validator       domain.ValidatorUsecase
	favoriteUsecase domain.FavoriteUsecase
}

func NewFavoriteController(
	validator domain.ValidatorUsecase,
	favoriteUsecase domain.FavoriteUsecase,
) *favoriteController {
	return &favoriteController{
		validator:       validator,
		favoriteUsecase: favoriteUsecase,
	}
}

func (fc *favoriteController) GetAll(ctx *fiber.Ctx) error {
	favorites, err := fc.favoriteUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    favorites,
	})
}

// GetByUserID godoc
// @Summary Get favorite by user id
// @Description Get favorite by user id
// @Tags Favorite
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} domain.Response{data=[]domain.Favorite}
// @Router /api/favorite [get]
func (fc *favoriteController) GetByUserID(ctx *fiber.Ctx) error {
	userID := ctx.Locals(constant.USER_ID).(string)
	favorites, err := fc.favoriteUsecase.GetByUserID(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    favorites,
	})
}

// Create godoc
// @Summary Create favorite
// @Description Create favorite
// @Tags Favorite
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body payload.CreateFavoriteDTO true "Favorite"
// @Success 201 {object} domain.Response{data=[]domain.Favorite}
// @Router /api/favorite [post]
func (fc *favoriteController) Create(ctx *fiber.Ctx) error {
	var body payload.CreateFavoriteDTO
	if err := fc.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	userID := ctx.Locals(constant.USER_ID).(string)
	favoriteCreate := domain.Favorite{
		USER_ID:    userID,
		PRODUCT_ID: body.PRODUCT_ID,
	}
	if err := fc.favoriteUsecase.Create(&favoriteCreate); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	favorite, err := fc.favoriteUsecase.GetByUserID(userID)
	if err != nil {
		return ctx.Status(fiber.StatusCreated).JSON(&domain.Response{
			MESSAGE: constant.MESSAGE_SUCCESS,
			SUCCESS: true,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    favorite,
	})
}

// Delete godoc
// @Summary Delete favorite
// @Description Delete favorite
// @Tags Favorite
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Favorite ID"
// @Success 200 {object} domain.Response
// @Router /api/favorite/{id} [delete]
func (fc *favoriteController) Delete(ctx *fiber.Ctx) error {
	favoriteID := ctx.Params("id")
	if favoriteID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&domain.Response{
			MESSAGE: constant.MESSAGE_NOT_FOUND,
			SUCCESS: false,
		})
	}

	userID := ctx.Locals(constant.USER_ID).(string)
	if err := fc.favoriteUsecase.Delete(userID, favoriteID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
	})
}
