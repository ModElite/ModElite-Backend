package usecase

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
)

type productUsecase struct {
	productRepo domain.ProductRepository
}

func NewProductUsecase(
	productRepo domain.ProductRepository,
) domain.ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
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

	return products, nil
}

func (u *productUsecase) GetProductWithOptionsAndSizes(productId string) (*domain.Product, error) {
	product, err := u.productRepo.GetProductWithOptionsAndSizes(productId)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	return product, nil
}

func (u *productUsecase) GetProductsBySeller(sellerID string) (*[]domain.Product, error) {
	products, err := u.productRepo.GetProductsBySeller(sellerID)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
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
