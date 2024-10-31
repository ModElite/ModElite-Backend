package usecase

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
)

type productUsecase struct {
	productRepo          domain.ProductRepository
	productOptionUsecase domain.ProductOptionUsecase
	productSizeUsecase   domain.ProductSizeUsecase
	sizeUsecase          domain.SizeUsecase
}

func NewProductUsecase(
	productRepo domain.ProductRepository,
	productOptionUsecase domain.ProductOptionUsecase,
	productSizeUsecase domain.ProductSizeUsecase,
	sizeUsecase domain.SizeUsecase,
) domain.ProductUsecase {
	return &productUsecase{
		productRepo:          productRepo,
		productOptionUsecase: productOptionUsecase,
		productSizeUsecase:   productSizeUsecase,
		sizeUsecase:          sizeUsecase,
	}
}

func (u *productUsecase) GetAll() (*[]domain.Product, error) {
	products, err := u.productRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	for i := range *products {
		product := &(*products)[i]
		productOptions, err := u.productOptionUsecase.GetByProductID(product.ID)
		if err != nil {
			return nil, fmt.Errorf("error product getall: %w", err)
		}
		product.PRODUCT_OPTION = productOptions

		for j := range *product.PRODUCT_OPTION {
			productOption := &(*product.PRODUCT_OPTION)[j]
			productSizes, err := u.productSizeUsecase.GetByProductOptionID(productOption.ID)
			if err != nil {
				return nil, fmt.Errorf("error product getall: %w", err)
			}

			for k := range *productSizes {
				productSize := &(*productSizes)[k]
				size, err := u.sizeUsecase.GetByID(productSize.SIZE_ID)
				if err != nil {
					return nil, fmt.Errorf("error product getall: %w", err)
				}
				productSize.SIZE = size
			}

			productOption.PRODUCT_SIZE = productSizes
		}
	}
	return products, nil
}

func (u *productUsecase) GetByID(id string) (*domain.Product, error) {
	product, err := u.productRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	} else if product == nil {
		return nil, fmt.Errorf(constant.MESSAGE_NOT_FOUND)
	}

	productOptions, err := u.productOptionUsecase.GetByProductID(product.ID)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}
	product.PRODUCT_OPTION = productOptions

	for j := range *product.PRODUCT_OPTION {
		productOption := &(*product.PRODUCT_OPTION)[j]
		productSizes, err := u.productSizeUsecase.GetByProductOptionID(productOption.ID)
		if err != nil {
			return nil, fmt.Errorf("error product getall: %w", err)
		}

		for k := range *productSizes {
			productSize := &(*productSizes)[k]
			size, err := u.sizeUsecase.GetByID(productSize.SIZE_ID)
			if err != nil {
				return nil, fmt.Errorf("error product getall: %w", err)
			}
			productSize.SIZE = size
		}

		productOption.PRODUCT_SIZE = productSizes
	}

	return product, nil
}

func (u *productUsecase) GetBySellerID(SellerID string) (*[]domain.Product, error) {
	products, err := u.productRepo.GetBySellerID(SellerID)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	for i := range *products {
		product := &(*products)[i]
		productOptions, err := u.productOptionUsecase.GetByProductID(product.ID)
		if err != nil {
			return nil, fmt.Errorf("error product getall: %w", err)
		}
		product.PRODUCT_OPTION = productOptions

		for j := range *product.PRODUCT_OPTION {
			productOption := &(*product.PRODUCT_OPTION)[j]
			productSizes, err := u.productSizeUsecase.GetByProductOptionID(productOption.ID)
			if err != nil {
				return nil, fmt.Errorf("error product getall: %w", err)
			}

			for k := range *productSizes {
				productSize := &(*productSizes)[k]
				size, err := u.sizeUsecase.GetByID(productSize.SIZE_ID)
				if err != nil {
					return nil, fmt.Errorf("error product getall: %w", err)
				}
				productSize.SIZE = size
			}

			productOption.PRODUCT_SIZE = productSizes
		}
	}
	return products, nil
}
