package usecase

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/google/uuid"
)

type sellerUsecase struct {
	sellerRepo  domain.SellerRepository
	userUsecase domain.UserUsecase
}

func NewSellerUsecase(
	sellerRepo domain.SellerRepository,
	userUsecase domain.UserUsecase,
) domain.SellerUsecase {
	return &sellerUsecase{
		sellerRepo:  sellerRepo,
		userUsecase: userUsecase,
	}
}

func (u *sellerUsecase) GetAll(userId string) (*[]domain.Seller, error) {
	Permission, err := u.userUsecase.CheckAdmin(userId)
	if err != nil {
		return nil, err
	} else if !Permission {
		return nil, fmt.Errorf(constant.MESSAGE_PERMISSION_DENIED)
	}

	sellers, err := u.sellerRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return sellers, nil
}

func (u *sellerUsecase) Create(data *domain.Seller) error {
	seller := &domain.Seller{
		ID:          uuid.New().String(),
		NAME:        data.NAME,
		DESCRIPTION: data.DESCRIPTION,
		LOGO_URL:    data.LOGO_URL,
		LOCATION:    data.LOCATION,
		OWNER_ID:    data.OWNER_ID,
		IS_VERIFY:   false,
	}

	err := u.sellerRepo.Create(seller)
	if err != nil {
		return err
	}
	return nil
}
