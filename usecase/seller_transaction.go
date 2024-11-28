package usecase

import "github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"

type sellerTransactionUsecase struct {
	sellerTransactionRepo domain.SellerTransactionRepository
}

func NewSellerTransactionUsecase(sellerTransactionRepo domain.SellerTransactionRepository) domain.SellerTransactionUsecase {
	return &sellerTransactionUsecase{
		sellerTransactionRepo: sellerTransactionRepo,
	}
}

func (u *sellerTransactionUsecase) GetAll() (*[]domain.SellerTransaction, error) {
	return u.sellerTransactionRepo.GetAll()
}

func (u *sellerTransactionUsecase) GetBySellerId(sellerId string) (*[]domain.SellerTransaction, error) {
	return u.sellerTransactionRepo.GetBySellerId(sellerId)
}
