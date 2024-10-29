package usecase

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
)

type productUsecase struct {
	productRepo   domain.ProductRepository
	sellerUsecase domain.SellerUsecase
	userUsecase   domain.UserUsecase
}

func NewProductUsecase(
	productRepo domain.ProductRepository,
	sellerUsecase domain.SellerUsecase,
	userUsecase domain.UserUsecase,
) domain.ProductUsecase {
	return &productUsecase{
		productRepo:   productRepo,
		sellerUsecase: sellerUsecase,
		userUsecase:   userUsecase,
	}
}

func (u *productUsecase) GetAll() (*[]domain.Product, error) {
	products, err := u.productRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (u *productUsecase) GetByID(id string) (*domain.Product, error) {
	product, err := u.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	} else if product == nil {
		return nil, fmt.Errorf(constant.MESSAGE_NOT_FOUND)
	}
	return product, nil
}

func (u *productUsecase) GetBySellerID(SellerID string) (*[]domain.Product, error) {
	products, err := u.productRepo.GetBySellerID(SellerID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
