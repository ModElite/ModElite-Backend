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
	for i := range order {
		orderProducts := make([]domain.OrderProductResponse, 0)
		err = r.db.Select(&orderProducts, `
		SELECT
			order_product."id" AS "id", 
			order_product.order_id AS order_id, 
			order_product.product_size_id AS product_size_id,
			order_product.quantity AS quantity, 
			order_product.price AS price, 
			order_product.created_at AS created_at, 
			order_product.updated_at AS updated_at, 
			"size"."size" AS "size", 
			product_option.label AS label, 
			product_option.image_url AS image_url, 
			product."name" AS "name", 
			product.description AS description, 
			product.price AS product_price, 
			product."image_url" AS "product_image_url", 
			seller."name" AS seller_name, 
			seller.logo_url AS seller_logo_url, 
			seller."id" AS seller_id
		FROM
			order_product
			INNER JOIN
			product_size
			ON 
				order_product.product_size_id = product_size."id"
			INNER JOIN
			product_option
			ON 
				product_size.product_option_id = product_option."id"
			INNER JOIN
			product
			ON 
				product_option.product_id = product."id"
			INNER JOIN
			"size"
			ON 
				product_size.size_id = "size"."id"
			INNER JOIN
			seller
			ON 
				product.seller_id = seller."id"
		WHERE order_id = $1;
	`, order[i].ID)
		if err != nil {
			return nil, err
		}
		order[i].ORDER_PRODUCT_DATA = &orderProducts
	}
	return &order, nil
}

func (r *orderRepository) CreateOrder(order *[]domain.OrderProduct, address string, voucherId *string, shipping_price float64, totalPrice float64, toDiscount float64, userId string, firstName string, lastName string, email string, phone string) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	id := uuid.New().String()
	if *voucherId == "" {
		voucherId = nil
	}

	_, err = tx.Exec(`INSERT INTO "order" (id, user_id, status, total_price, product_price, shipping_price, discount, voucher_code, address, first_name, last_name, email, phone, express_provider, express_tracking_number) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, '', '');`,
		id, userId, domain.ORDER_PENDING, totalPrice+shipping_price, totalPrice, shipping_price, toDiscount, voucherId, address, firstName, lastName, email, phone)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return "", err
		}
		return "", err
	}
	// INSERT INTO order_product (order_id, product_size_id, quantity, price, status) VALUES ($1, $2, $3, $4, $5);
	for _, orderProduct := range *order {
		orderProductId := uuid.New().String()
		_, err = tx.Exec(`INSERT INTO order_product (id ,order_id, product_size_id, quantity, price) VALUES ($1, $2, $3, $4, $5); `,
			orderProductId, id, orderProduct.PRODUCT_SIZE_ID, orderProduct.QUANTITY, orderProduct.PRICE)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return "", err
			}
			return "", err
		}
		_, err = tx.Exec(`UPDATE product_size SET quantity = quantity - $1 WHERE id = $2;`, orderProduct.QUANTITY, orderProduct.PRODUCT_SIZE_ID)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return "", err
			}
			return "", err
		}
	}
	return id, tx.Commit()
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
	orderProducts := make([]domain.OrderProductResponse, 0)
	err = r.db.Select(&orderProducts, `
		SELECT
			order_product."id" AS "id", 
			order_product.order_id AS order_id, 
			order_product.product_size_id AS product_size_id,
			order_product.quantity AS quantity, 
			order_product.price AS price, 
			order_product.created_at AS created_at, 
			order_product.updated_at AS updated_at, 
			"size"."size" AS "size", 
			product_option.label AS label, 
			product_option.image_url AS image_url, 
			product."name" AS "name", 
			product.description AS description, 
			product.price AS product_price, 
			product."image_url" AS "product_image_url", 
			seller."name" AS seller_name, 
			seller.logo_url AS seller_logo_url, 
			seller."id" AS seller_id
		FROM
			order_product
			INNER JOIN
			product_size
			ON 
				order_product.product_size_id = product_size."id"
			INNER JOIN
			product_option
			ON 
				product_size.product_option_id = product_option."id"
			INNER JOIN
			product
			ON 
				product_option.product_id = product."id"
			INNER JOIN
			"size"
			ON 
				product_size.size_id = "size"."id"
			INNER JOIN
			seller
			ON 
				product.seller_id = seller."id"
		WHERE order_id = $1;
	`, order.ID)
	if err != nil {
		return nil, err
	}
	order.ORDER_PRODUCT_DATA = &orderProducts

	return &order, nil
}

func (r *orderRepository) CheckSellerUserID(SellerID string, UserID string) (bool, error) {
	var count int
	err := r.db.Get(&count, `SELECT COUNT(*) FROM seller WHERE id = $1 AND owner_id = $2;`, SellerID, UserID)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (r *orderRepository) GetSellerOrder(SellerID string) (*[]domain.Order, error) {
	order := make([]domain.Order, 0)
	err := r.db.Select(&order, `
		SELECT
			"order".*
		FROM
			"order"
		WHERE
			"order"."id" IN (
			SELECT
				order_product.order_id 
			FROM
				order_product
				INNER JOIN product_size ON order_product.product_size_id = product_size."id"
				INNER JOIN product_option ON product_size.product_option_id = product_option."id"
				INNER JOIN product ON product_option.product_id = product."id" 
			WHERE
				product.seller_id = $1 
			GROUP BY
			order_product.order_id 
			)`, SellerID)
	if err != nil {
		return nil, err
	}
	for i := range order {
		orderProducts := make([]domain.OrderProductResponse, 0)
		err = r.db.Select(&orderProducts, `
		SELECT
			order_product."id" AS "id", 
			order_product.order_id AS order_id, 
			order_product.product_size_id AS product_size_id, 
			order_product.quantity AS quantity, 
			order_product.price AS price, 
			order_product.created_at AS created_at, 
			order_product.updated_at AS updated_at, 
			"size"."size" AS "size", 
			product_option.label AS label, 
			product_option.image_url AS image_url, 
			product."name" AS "name", 
			product.description AS description, 
			product.price AS product_price, 
			product."image_url" AS "product_image_url", 
			seller."name" AS seller_name, 
			seller.logo_url AS seller_logo_url, 
			seller."id" AS seller_id
		FROM
			order_product
			INNER JOIN
			product_size
			ON 
				order_product.product_size_id = product_size."id"
			INNER JOIN
			product_option
			ON 
				product_size.product_option_id = product_option."id"
			INNER JOIN
			product
			ON 
				product_option.product_id = product."id"
			INNER JOIN
			"size"
			ON 
				product_size.size_id = "size"."id"
			INNER JOIN
			seller
			ON 
				product.seller_id = seller."id"
		WHERE order_id = $1;
	`, order[i].ID)
		if err != nil {
			return nil, err
		}
		order[i].ORDER_PRODUCT_DATA = &orderProducts
	}
	return &order, nil
}

func (r *orderRepository) UpdateOrderExpress(orderID string, expressProvider string, expressTrackingNumber string) error {
	_, err := r.db.Exec(`UPDATE "order" SET express_provider = $1, express_tracking_number = $2, "status" = 'DELIVERY_ON_THE_WAY'  WHERE id = $3;`, expressProvider, expressTrackingNumber, orderID)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) FakePayment(orderID string) error {
	_, err := r.db.Exec(`UPDATE "order" SET status = 'PAYMENT_SUCCESS' WHERE id = $1;`, orderID)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) GetOrderPaymentDetail(orderID string) (*domain.OrderPaymentResponse, error) {
	var data domain.OrderPaymentResponse
	err := r.db.Get(&data, `
		SELECT
			id AS "order_id",
			total_price as "amount"
		FROM
			"order"
		WHERE
			id = $1;
	`, orderID)

	if err != nil {
		return nil, err
	}
	return &data, nil
}
