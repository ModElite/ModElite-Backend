package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type productController struct {
	validator      domain.ValidatorUsecase
	productUseCase domain.ProductUsecase
	sellerUseCase  domain.SellerUsecase
}

func NewProductController(validator domain.ValidatorUsecase,
	productUseCase domain.ProductUsecase,
	sellerUseCase domain.SellerUsecase,
) *productController {
	return &productController{
		validator:      validator,
		productUseCase: productUseCase,
		sellerUseCase:  sellerUseCase,
	}
}

// @Summary Get all products
// @Tags Product
// @Produce json
// @Success 200 {object} domain.Response
// @Router /api/product [get]
func (p *productController) GetAll(ctx *fiber.Ctx) error {
	products, err := p.productUseCase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    products,
	})
}

// @Summary Get products by seller id
// @Tags Product
// @Produce json
// @Param id path string true "Seller ID"
// @Success 200 {object} domain.Response
// @Router /api/product/seller/{id} [get]
func (p *productController) GetBySellerID(ctx *fiber.Ctx) error {
	if seller, err := p.sellerUseCase.GetByID(ctx.Params("id")); err == nil && seller == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	} else if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	products, err := p.productUseCase.GetBySellerID(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    products,
	})

}

// @Summary Get product by id
// @Tags Product
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} domain.Response
// @Router /api/product/{id} [get]
func (p *productController) GetByID(ctx *fiber.Ctx) error {
	product, err := p.productUseCase.GetByID(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	} else if product == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    product,
	})
}
