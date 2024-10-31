package usecase

import "github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"

type productSizeUsecase struct {
	productSizeRepo domain.ProductSizeRepository
}

func NewProductSizeUsecase(productSizeRepo domain.ProductSizeRepository) domain.ProductSizeUsecase {
	return &productSizeUsecase{
		productSizeRepo: productSizeRepo,
	}
}

func (pu *productSizeUsecase) GetAll() (*[]domain.ProductSize, error) {
	return pu.productSizeRepo.GetAll()
}

func (pu *productSizeUsecase) GetByID(id string) (*domain.ProductSize, error) {
	return pu.productSizeRepo.GetByID(id)
}

func (pu *productSizeUsecase) GetByProductOptionID(productOptionID string) (*[]domain.ProductSize, error) {
	return pu.productSizeRepo.GetByProductOptionID(productOptionID)
}

func (pu *productSizeUsecase) Create(productSize *domain.ProductSize) error {
	return pu.productSizeRepo.Create(productSize)
}

func (pu *productSizeUsecase) Update(productSize *domain.ProductSize) error {
	return pu.productSizeRepo.Update(productSize)
}
