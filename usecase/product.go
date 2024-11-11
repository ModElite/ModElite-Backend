package usecase

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/google/uuid"
)

type productUsecase struct {
	productRepo          domain.ProductRepository
	productOptionUsecase domain.ProductOptionUsecase
	productSizeUsecase   domain.ProductSizeUsecase
	tagUsecase           domain.TagUsecase
}

func NewProductUsecase(
	productRepo domain.ProductRepository,
	productOptionUsecase domain.ProductOptionUsecase,
	productSizeUsecase domain.ProductSizeUsecase,
	tagUsecase domain.TagUsecase,
) domain.ProductUsecase {
	return &productUsecase{
		productRepo:          productRepo,
		productOptionUsecase: productOptionUsecase,
		productSizeUsecase:   productSizeUsecase,
		tagUsecase:           tagUsecase,
	}
}

func (u *productUsecase) GetAll() (*[]domain.Product, error) {
	products, err := u.productRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	return products, nil
}

func (u *productUsecase) GetAllProductWithOptionsAndSizes() (*[]domain.Product, error) {
	products, err := u.productRepo.GetAllProductWithOptionsAndSizes()
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	for i, product := range *products {
		productTags, err := u.tagUsecase.GetTagByProductID(product.ID)
		if err != nil {
			return nil, fmt.Errorf("error product getall: %w", err)
		}
		(*products)[i].TAGS = productTags
	}

	return products, nil
}

func (u *productUsecase) GetProductWithOptionsAndSizes(productId string) (*domain.Product, error) {
	product, err := u.productRepo.GetProductWithOptionsAndSizes(productId)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	productTags, err := u.tagUsecase.GetTagByProductID(product.ID)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}
	product.TAGS = productTags

	return product, nil
}

func (u *productUsecase) GetProductsBySeller(sellerID string) (*[]domain.Product, error) {
	products, err := u.productRepo.GetProductsBySeller(sellerID)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	for i, product := range *products {
		productTags, err := u.tagUsecase.GetTagByProductID(product.ID)
		if err != nil {
			return nil, fmt.Errorf("error product getall: %w", err)
		}
		(*products)[i].TAGS = productTags
	}

	return products, nil
}

func (u *productUsecase) GetByID(id string) (*domain.Product, error) {
	product, err := u.productRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	return product, nil
}

func (u *productUsecase) GetBySellerID(SellerID string) (*[]domain.Product, error) {
	products, err := u.productRepo.GetBySellerID(SellerID)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	return products, nil
}

func (u *productUsecase) Create(product *domain.Product) (*string, error) {
	product.ID = uuid.New().String()

	if err := u.productRepo.Create(&domain.Product{
		ID:          product.ID,
		SELLER_ID:   product.SELLER_ID,
		NAME:        product.NAME,
		DESCRIPTION: product.DESCRIPTION,
		PRICE:       product.PRICE,
		IMAGE_URL:   product.IMAGE_URL,
		STATUS:      string(domain.ProductActive),
	}); err != nil {
		return nil, fmt.Errorf("error product create: %w", err)
	}

	for _, option := range *product.PRODUCT_OPTION {
		option.ID = uuid.New().String()
		option.PRODUCT_ID = product.ID
		if err := u.productOptionUsecase.Create(&domain.ProductOption{
			ID:         option.ID,
			PRODUCT_ID: option.PRODUCT_ID,
			LABEL:      option.LABEL,
			IMAGE_URL:  option.IMAGE_URL,
		}); err != nil {
			return nil, fmt.Errorf("error product option create: %w", err)
		}

		for _, size := range *option.PRODUCT_SIZE {
			size.ID = uuid.New().String()
			size.PRODUCT_OPTION_ID = option.ID
			if err := u.productSizeUsecase.Create(&domain.ProductSize{
				ID:                size.ID,
				PRODUCT_OPTION_ID: size.PRODUCT_OPTION_ID,
				SIZE_ID:           size.SIZE_ID,
				QUANTITY:          size.QUANTITY,
			}); err != nil {
				return nil, fmt.Errorf("error product size create: %w", err)
			}
		}
	}

	return &product.ID, nil
}
