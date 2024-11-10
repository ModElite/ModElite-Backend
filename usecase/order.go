package usecase

import "github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"

type orderUsecase struct {
	orderRepo domain.OrderRepository
}

func NewOrderUsecase(orderRepo domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo: orderRepo,
	}
}

func (u *orderUsecase) GetAll() (*[]domain.Order, error) {
	return u.orderRepo.GetAll()
}

func (u *orderUsecase) GetSelfOrder(userID string) (*[]domain.Order, error) {
	return u.orderRepo.GetSelfOrder(userID)
}
