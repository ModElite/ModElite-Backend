package repository

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type sellerTransactionRepository struct {
	db *sqlx.DB
}

func NewSellerTransactionRepository(db *sqlx.DB) domain.SellerTransactionRepository {
	return &sellerTransactionRepository{
		db: db,
	}
}

func (r *sellerTransactionRepository) GetAll() (*[]domain.SellerTransaction, error) {
	sellerTransactions := make([]domain.SellerTransaction, 0)
	err := r.db.Select(&sellerTransactions, "SELECT * FROM seller_transaction")
	if err != nil {
		return nil, err
	}

	return &sellerTransactions, nil
}

func (r *sellerTransactionRepository) GetBySellerId(sellerId string) (*[]domain.SellerTransaction, error) {
	sellerTransactions := make([]domain.SellerTransaction, 0)
	err := r.db.Select(&sellerTransactions, "SELECT * FROM seller_transaction WHERE seller_id = $1", sellerId)
	if err != nil {
		return nil, err
	}

	return &sellerTransactions, nil
}
