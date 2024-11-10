package usecase

import "github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"

type voucherUsecase struct {
	voucherRepo domain.VoucherRepository
}

func NewVoucherUsecase(voucherRepo domain.VoucherRepository) domain.VoucherUsecase {
	return &voucherUsecase{
		voucherRepo: voucherRepo,
	}
}

func (v *voucherUsecase) GetByID(id string) (*domain.Voucher, error) {
	return v.voucherRepo.GetByID(id)
}
