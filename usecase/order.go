package usecase

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
)

type orderUsecase struct {
	orderRepo   domain.OrderRepository
	productRepo domain.ProductRepository
}

func NewOrderUsecase(orderRepo domain.OrderRepository, productRepo domain.ProductRepository) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (u *orderUsecase) GetAll() (*[]domain.Order, error) {
	return u.orderRepo.GetAll()
}

func (u *orderUsecase) GetSelfOrder(userID string) (*[]domain.Order, error) {
	return u.orderRepo.GetSelfOrder(userID)
}

// c.orderUsecase.CreateOrder(&orderProdcts, address, payload.VOUCHER_ID, totalPrice, toDiscount, userID)
func (u *orderUsecase) CreateOrder(order *[]domain.OrderProduct, address string, voucherId *string, shipping_price float64, totalPrice float64, toDiscount float64, userId string, firstName string, lastName string, email string, phone string) (string, error) {
	return u.orderRepo.CreateOrder(order, address, voucherId, shipping_price, totalPrice, toDiscount, userId, firstName, lastName, email, phone)
}

func (u *orderUsecase) GetProductDetail(productSizeID string, quantity int) (*domain.OrderProduct, error) {
	// Get Price and Quantity from ProductSize
	quantityData, err := u.productRepo.GetProductPriceQuantity(productSizeID)
	if err != nil {
		return nil, err
	}
	// Make quantityData to OrderProduct
	orderProduct := domain.OrderProduct{
		PRODUCT_SIZE_ID: productSizeID,
		QUANTITY:        quantity,
		PRICE:           quantityData.Price,
		SELLER_ID:       quantityData.SellerID,
	}
	return &orderProduct, nil
}

func (u *orderUsecase) GetSelfOrderDetail(orderID string, userID string) (*domain.Order, error) {
	return u.orderRepo.GetSelfOrderDetail(orderID, userID)
}

func (u *orderUsecase) GetSellerOrder(SellerID string) (*[]domain.Order, error) {
	// Check Seller ID and UserID is the same
	order, err := u.orderRepo.GetSellerOrder(SellerID)
	// Add username to order data
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (u *orderUsecase) CheckSellerUserID(SellerID string, UserID string) error {
	isSame, err := u.orderRepo.CheckSellerUserID(SellerID, UserID)
	if err != nil {
		return err
	}
	if !isSame {
		return fmt.Errorf("UserID or SellerID invalid")
	}
	return nil
}

func (u *orderUsecase) UpdateOrderExpress(orderID string, expressProvider string, expressTrackingNumber string) error {
	return u.orderRepo.UpdateOrderExpress(orderID, expressProvider, expressTrackingNumber)
}

func (u *orderUsecase) FakePayment(orderID string) error {
	return u.orderRepo.FakePayment(orderID)
}

func (u *orderUsecase) GetOrderPaymentDetail(orderID string) (*domain.OrderPaymentResponse, error) {
	return u.orderRepo.GetOrderPaymentDetail(orderID)
}