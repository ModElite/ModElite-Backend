package usecase

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
)

type addressUsecase struct {
	addressRepo domain.AddressRepository
	userUsecase domain.UserUsecase
}

func NewAddressUsecase(
	addressRepo domain.AddressRepository,
	userUsecase domain.UserUsecase,
) domain.AddressUsecase {
	return &addressUsecase{
		addressRepo: addressRepo,
		userUsecase: userUsecase,
	}
}

func (a *addressUsecase) CheckPermissionCanModifyAddress(userID string, addressID int) (bool, error) {
	addresse, err := a.addressRepo.GetById(addressID)
	if err != nil || addresse == nil {
		return false, err
	} else if addresse.USER_ID != userID {
		if isAdmin, err := a.userUsecase.CheckAdmin(userID); err != nil {
			return false, nil
		} else if !isAdmin {
			return false, nil
		}
	}

	return true, nil
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
	id, err := a.addressRepo.Create(address)
	if err != nil {
		return err
	}

	if address.DEFAULT {
		if err = a.addressRepo.UpdateDefaultByUserId(address.USER_ID, id); err != nil {
			return err
		}
	}
	return err
}

func (a *addressUsecase) Update(address *domain.Address) error {
	err := a.addressRepo.Update(address)
	if err != nil {
		return err
	}

	if address.DEFAULT {
		if address.USER_ID == "" {
			return fmt.Errorf("user id is empty")
		}
		err = a.addressRepo.UpdateDefaultByUserId(address.USER_ID, address.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *addressUsecase) Delete(addressID int) error {
	return a.addressRepo.Delete(addressID)
}

func (a *addressUsecase) AddressIdToString(addressId int) (string, error) {
	address, err := a.addressRepo.GetById(addressId)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Name: %s %s Phone: %s Email: %s Address: %s %s %s %s %s",
		address.FIRST_NAME,
		address.LAST_NAME,
		address.EMAIL,
		address.PHONE,
		address.ADDRESS,
		address.SUB_DISTRICT,
		address.DISTRICT,
		address.PROVINCE,
		address.ZIP_CODE,
	), nil
}
