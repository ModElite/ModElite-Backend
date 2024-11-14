package usecase

import "github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"

type productOptionUsecase struct {
	productOptionRepo domain.ProductOptionRepository
	productSizeRepo   domain.ProductSizeRepository
}

func NewProductOptionUsecase(
	productOptionRepo domain.ProductOptionRepository,
	productSizeRepo domain.ProductSizeRepository,
) domain.ProductOptionUsecase {
	return &productOptionUsecase{
		productOptionRepo: productOptionRepo,
		productSizeRepo:   productSizeRepo,
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

func (pu *productOptionUsecase) GetByProductIDAndFilterActive(productID string) (*[]domain.ProductOption, error) {
	return pu.productOptionRepo.GetByProductIDAndFilterActive(productID)
}

func (pu *productOptionUsecase) Create(productOption *domain.ProductOption) error {
	return pu.productOptionRepo.Create(productOption)
}

func (pu *productOptionUsecase) Update(productOption *domain.ProductOption) error {
	return pu.productOptionRepo.Update(productOption)
}

func (pu *productOptionUsecase) SoftDeleteProductOptionAndSizeByProductID(ProductOptionId string) error {
	if err := pu.productOptionRepo.SoftDelete(ProductOptionId); err != nil {
		return err
	}

	if err := pu.productSizeRepo.SoftDeleteByProductOptionID(ProductOptionId); err != nil {
		return err
	}

	return nil
}
