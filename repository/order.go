package repository

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) domain.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) GetAll() (*[]domain.Order, error) {
	orders := make([]domain.Order, 0)
	err := r.db.Select(&orders, `SELECT * FROM "order";`)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (r *orderRepository) GetSelfOrder(userID string) (*[]domain.Order, error) {
	order := make([]domain.Order, 0)
	err := r.db.Select(&order, `SELECT * FROM "order" WHERE user_id = $1;`, userID)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
