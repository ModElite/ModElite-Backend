package usecase

import "github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"

type productOptionUsecase struct {
	productOptionRepo domain.ProductOptionRepository
}

func NewProductOptionUsecase(productOptionRepo domain.ProductOptionRepository) domain.ProductOptionUsecase {
	return &productOptionUsecase{
		productOptionRepo: productOptionRepo,
	}
}

func (pu *productOptionUsecase) GetAll() (*[]domain.ProductOption, error) {
	return pu.productOptionRepo.GetAll()
}

func (pu *productOptionUsecase) GetByID(id string) (*domain.ProductOption, error) {
	return pu.productOptionRepo.GetByID(id)
}

func (pu *productOptionUsecase) GetByProductID(productID string) (*[]domain.ProductOption, error) {
	return pu.productOptionRepo.GetByProductID(productID)
}

func (pu *productOptionUsecase) Create(productOption *domain.ProductOption) error {
	return pu.productOptionRepo.Create(productOption)
}

func (pu *productOptionUsecase) Update(productOption *domain.ProductOption) error {
	return pu.productOptionRepo.Update(productOption)
}
