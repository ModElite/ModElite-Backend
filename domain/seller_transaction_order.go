package domain

type SellerTransactionOrder struct {
	SELLER_TRANSACTION_ID string `json:"sellerTransactionId" db:"seller_transaction_id"`
	ORDER_ID              string `json:"orderId" db:"order_id"`
	CREATED_AT            string `json:"createdAt" db:"created_at"`
	UPDATED_AT            string `json:"updatedAt" db:"updated_at"`
}

type SellerTransactionOrderRepository interface {

}

type SellerTransactionOrderUsecase interface {
}
