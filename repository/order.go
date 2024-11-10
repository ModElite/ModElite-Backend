package repository

import (
	"database/sql"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/google/uuid"
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

func (r *orderRepository) CreateOrder(order *[]domain.OrderProduct, address string, voucherId *string, shipping_price float64, totalPrice float64, toDiscount float64, userId string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	id := uuid.New().String()
	if *voucherId == "" {
		voucherId = nil
	}

	_, err = tx.Exec(`INSERT INTO "order" (id, user_id, status, total_price, product_price, shipping_price, discount, voucher_code, address) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`,
		id, userId, domain.ORDER_PENDING, totalPrice+shipping_price, totalPrice, shipping_price, toDiscount, voucherId, address)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	// INSERT INTO order_product (order_id, product_size_id, quantity, price, status) VALUES ($1, $2, $3, $4, $5);
	for _, orderProduct := range *order {
		orderProductId := uuid.New().String()
		_, err = tx.Exec(`INSERT INTO order_product (id ,order_id, product_size_id, quantity, price, status) VALUES ($1, $2, $3, $4, $5, $6); UPDATE product_size SET quantity = quantity - $4 WHERE id = $3;`,
			orderProductId, id, orderProduct.PRODUCT_SIZE_ID, orderProduct.QUANTITY, orderProduct.PRICE, domain.ORDER_PRODUCT_PENDING)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}
	}
	return tx.Commit()
}

func (r *orderRepository) GetSelfOrderDetail(orderID string, userID string) (*domain.Order, error) {
	order := domain.Order{}
	err := r.db.Get(&order, `SELECT * FROM "order" WHERE id = $1 AND user_id = $2;`, orderID, userID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	orderProducts := make([]domain.OrderProduct, 0)
	err = r.db.Select(&orderProducts, `SELECT * FROM order_product WHERE order_id = $1;`, orderID)
	if err != nil {
		return nil, err
	}
	order.ORDER_PRODUCT = &orderProducts
	return &order, nil
}
