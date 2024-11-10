package usecase

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
)

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

func (v *voucherUsecase) Search(code string) (*domain.Voucher, error) {
	return v.voucherRepo.Search(code)
}

func (v *voucherUsecase) CreateVoucher(voucher *domain.Voucher) error {
	return v.voucherRepo.CreateVoucher(voucher)
}

func (v *voucherUsecase) CheckDuplicateCode(code string) bool {
	return v.voucherRepo.CheckDuplicateCode(code)
}
