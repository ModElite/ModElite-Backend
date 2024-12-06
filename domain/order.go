package domain

import "time"

type OrderStatusType string

const (
	ORDER_PENDING             OrderStatusType = "PENDING"
	ORDER_PAYMENT_SUCCESS     OrderStatusType = "PAYMENT_SUCCESS"
	ORDER_DELIVERY_ON_THE_WAY OrderStatusType = "DELIVERY_ON_THE_WAY"
	ORDER_REFUND              OrderStatusType = "REFUND"
	ORDER_END                 OrderStatusType = "END"
	ORDER_CANCEL              OrderStatusType = "CANCEL"
)

type Order struct {
	ID                             string                  `json:"id" db:"id"`
	ORDER_PRODUCT                  *[]OrderProduct         `json:"orderProduct,omitempty" db:"-"`
	ORDER_PRODUCT_DATA             *[]OrderProductResponse `json:"orderProductData,omitempty" db:"-"`
	USER                           *User                   `json:"user,omitempty" db:"-"`
	USER_ID                        string                  `json:"userId" db:"user_id"`
	FIRSTNAME                      string                  `json:"firstName,omitempty" db:"first_name"`
	LASTNAME                       string                  `json:"lastName,omitempty" db:"last_name"`
	EMAIL                          string                  `json:"email,omitempty" db:"email"`
	PHONE                          string                  `json:"phone,omitempty" db:"phone"`
	STATUS                         OrderStatusType         `json:"status" db:"status"`
	SELLER_PAYMENT_STATUS          bool                    `json:"sellerPaymentStatus" db:"seller_payment_status"`
	SELLER_PAYMENT_PRODUCT_AMOUNT  float64                 `json:"sellerPaymentProductAmount" db:"seller_payment_product_amount"`
	SELLER_PAYMENT_SHIPPING_AMOUNT float64                 `json:"sellerPaymentShippingAmount" db:"seller_payment_shipping_amount"`
	EXPRESS_PROVIDER               string                  `json:"expressProvider,omitempty" db:"express_provider"`
	EXPRESS_TRACKING_NUMBER        string                  `json:"expressTrackingNumber,omitempty" db:"express_tracking_number"`
	TOTAL_PRICE                    float64                 `json:"totalPrice" db:"total_price"`
	PRODUCT_PRICE                  float64                 `json:"productPrice" db:"product_price"`
	SHIPPING_PRICE                 float64                 `json:"shippingPrice" db:"shipping_price"`
	DISCOUNT                       float64                 `json:"discount" db:"discount"`
	VOUCHER_CODE                   *string                 `json:"voucherCode" db:"voucher_code"`
	ADDRESS                        string                  `json:"address" db:"address"`
	CREATED_AT                     time.Time               `json:"createdAt" db:"created_at"`
	UPDATED_AT                     time.Time               `json:"updatedAt" db:"updated_at"`
}

type OrderPaymentResponse struct {
	ORDER_ID string  `json:"orderId"`
	AMOUNT   float64 `json:"amount"`
}

type OrderProductResponse struct {
	OrderProduct
	PRODUCT_OPTION_LABEL     string  `json:"productOptionLabel" db:"label"`
	PRODUCT_OPTION_IMAGE_URL string  `json:"productOptionImageUrl" db:"image_url"`
	PRODUCT_NAME             string  `json:"productName" db:"name"`
	PRODUCT_DESCRIPTION      string  `json:"productDescription" db:"description"`
	PRODUCT_PRICE            float64 `json:"productPrice" db:"product_price"`
	SIZE                     string  `json:"productSize" db:"size"`
	SELLER_NAME              string  `json:"sellerName" db:"seller_name"`
	SELLER_LOGO_URL          string  `json:"sellerLogoUrl" db:"seller_logo_url"`
	SELLER_ID                string  `json:"sellerId" db:"seller_id"`
	PRODUCT_IMAGE_URL        string  `json:"productImageUrl" db:"product_image_url"`
}

type OrderRepository interface {
	GetAll() (*[]Order, error)
	GetSelfOrder(userID string) (*[]Order, error)
	CreateOrder(order *[]OrderProduct, address string, voucherId *string, shipping_price float64, totalPrice float64, toDiscount float64, userId string, firstName string, lastName string, email string, phone string) (string, error)
	GetSelfOrderDetail(orderID string, userID string) (*Order, error)
	GetSellerOrder(SellerID string) (*[]Order, error)
	CheckSellerUserID(SellerID string, UserID string) (bool, error)
	UpdateOrderExpress(orderID string, expressProvider string, expressTrackingNumber string) error
	FakePayment(orderID string) error
	GetOrderPaymentDetail(orderID string) (*OrderPaymentResponse, error)
}

type OrderUsecase interface {
	GetAll() (*[]Order, error)
	GetSelfOrder(userID string) (*[]Order, error)
	CreateOrder(order *[]OrderProduct, address string, voucherId *string, shipping_price float64, totalPrice float64, toDiscount float64, userId string, firstName string, lastName string, email string, phone string) (string, error)
	GetProductDetail(productSizeID string, quantity int) (*OrderProduct, error)
	GetSelfOrderDetail(orderID string, userID string) (*Order, error)
	GetSellerOrder(SellerID string) (*[]Order, error)
	UpdateOrderExpress(orderID string, expressProvider string, expressTrackingNumber string) error
	CheckSellerUserID(SellerID string, UserID string) error
	FakePayment(orderID string) error
	GetOrderPaymentDetail(orderID string) (*OrderPaymentResponse, error)
}
