package usecase

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
)

type addressUsecase struct {
	addressRepo domain.AddressRepository
}

func NewAddressUsecase(
	addressRepo domain.AddressRepository,
) domain.AddressUsecase {
	return &addressUsecase{
		addressRepo: addressRepo,
	}
}

func (a *addressUsecase) GetAll() (*[]domain.Address, error) {
	return a.addressRepo.GetAll()
}

func (a *addressUsecase) GetAddressByID(addressID int) (*domain.Address, error) {
	return a.addressRepo.GetById(addressID)
}

func (a *addressUsecase) GetAddressByUserID(userID string) (*[]domain.Address, error) {
	return a.addressRepo.GetByUserId(userID)
}

func (a *addressUsecase) Create(address *domain.Address) error {
	return a.addressRepo.Create(address)
}

func (a *addressUsecase) Update(address *domain.Address) error {
	return a.addressRepo.Update(address)
}

func (a *addressUsecase) Delete(addressID int) error {
	return a.addressRepo.Delete(addressID)
}
