package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/payload"
	"github.com/gofiber/fiber/v2"
)

type productController struct {
	validator      domain.ValidatorUsecase
	productUseCase domain.ProductUsecase
	sellerUseCase  domain.SellerUsecase
	userUseCase    domain.UserUsecase
	tagUseCase     domain.TagUsecase
}

func NewProductController(
	validator domain.ValidatorUsecase,
	productUseCase domain.ProductUsecase,
	sellerUseCase domain.SellerUsecase,
	userUseCase domain.UserUsecase,
	tagUseCase domain.TagUsecase,
) *productController {
	return &productController{
		validator:      validator,
		productUseCase: productUseCase,
		sellerUseCase:  sellerUseCase,
		userUseCase:    userUseCase,
		tagUseCase:     tagUseCase,
	}
}

// @Summary Get all products
// @Tags Product
// @Produce json
// @Success 200 {object} domain.Response{data=[]domain.Product}
// @Router /api/product [get]
func (p *productController) GetAllProductWithOptionsAndSizes(ctx *fiber.Ctx) error {
	products, err := p.productUseCase.GetAllProductWithOptionsAndSizes()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&domain.Response{
			SUCCESS: false,
			MESSAGE: err.Error(),
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
// @Success 200 {object} domain.Response{data=[]domain.Product}
// @Router /api/product/seller/{id} [get]
func (p *productController) GetBySellerID(ctx *fiber.Ctx) error {
	sellerId := ctx.Params("id")
	if sellerId == "" || !p.validator.ValidateUUID(sellerId) {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}
	if seller, err := p.sellerUseCase.GetByID(sellerId); err == nil && seller == nil {
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

	products, err := p.productUseCase.GetProductsBySeller(sellerId)
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
// @Success 200 {object} domain.Response{data=domain.Product}
// @Router /api/product/{id} [get]
func (p *productController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" || !p.validator.ValidateUUID(id) {
		return ctx.Status(fiber.StatusNotFound).JSON(&domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}
	product, err := p.productUseCase.GetProductWithOptionsAndSizes(ctx.Params("id"))
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

// @Summary Create product
// @Tags Product
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body payload.CreateProductDTO true "Product Body"
// @Success 200 {object} domain.Response
// @Router /api/product [post]
func (p *productController) Create(ctx *fiber.Ctx) error {
	var body payload.CreateProductDTO
	if err := p.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}
	if seller, err := p.sellerUseCase.GetByID(body.SELLER_ID); err != nil && seller == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	} else if seller.OWNER_ID != ctx.Locals(constant.USER_ID).(string) {
		if isAdmin, err := p.userUseCase.CheckAdmin(ctx.Locals(constant.USER_ID).(string)); err != nil || !isAdmin {
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: constant.MESSAGE_UNAUTHORIZED,
			})
		}
	}

	productOption := make([]domain.ProductOption, 0)
	for _, option := range *body.PRODUCT_OPTION {
		productSize := make([]domain.ProductSize, 0)
		for _, size := range *option.PRODUCT_SIZE {
			productSize = append(productSize, domain.ProductSize{
				SIZE_ID:  size.SIZE_ID,
				QUANTITY: size.QUANTITY,
			})
		}
		productOption = append(productOption, domain.ProductOption{
			LABEL:        option.LABEL,
			IMAGE_URL:    option.IMAGE_URL,
			PRODUCT_SIZE: &productSize,
		})
	}

	product := &domain.Product{
		SELLER_ID:      body.SELLER_ID,
		NAME:           body.NAME,
		DESCRIPTION:    body.DESCRIPTION,
		PRICE:          body.PRICE,
		IMAGE_URL:      body.IMAGE_URL,
		PRODUCT_OPTION: &productOption,
	}

	id, err := p.productUseCase.Create(product)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	if body.TAG_ID != nil {
		for _, tagID := range *body.TAG_ID {
			if err := p.tagUseCase.CreateProductTag(*id, tagID); err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
					SUCCESS: false,
					MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
				})
			}
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}

func (p *productController) Update(ctx *fiber.Ctx) error {
	// TODO: Implement Update product with options and sizes and tags
	return nil
}

// @Summary Delete product
// @Tags Product
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Success 200 {object} domain.Response
// @Router /api/product/{id} [delete]
func (p *productController) SoftDelete(ctx *fiber.Ctx) error {
	userId := ctx.Locals(constant.USER_ID).(string)
	productId := ctx.Params("id", "")
	if productId == "" || !p.validator.ValidateUUID(productId) {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}

	if canEdit, err := p.productUseCase.CheckPermissionCanModifyProduct(userId, productId); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	} else if !canEdit {
		return ctx.Status(fiber.StatusForbidden).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_PERMISSION_DENIED,
			SUCCESS: false,
		})
	}

	if err := p.productUseCase.SoftDeleteWithOptionsAndSizes(productId); err != nil {
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
