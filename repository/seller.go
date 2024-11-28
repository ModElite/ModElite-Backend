package repository

import (
	"database/sql"
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type sellerRepository struct {
	db *sqlx.DB
}

func NewSellerRepository(db *sqlx.DB) domain.SellerRepository {
	return &sellerRepository{
		db: db,
	}
}

func (r *sellerRepository) GetAll() (*[]domain.Seller, error) {
	sellers := make([]domain.Seller, 0)
	err := r.db.Select(&sellers, "SELECT * FROM seller")
	if err != nil {
		return nil, fmt.Errorf("error get all seller: %v", err)
	}
	return &sellers, nil
}

func (r *sellerRepository) GetByID(id string) (*domain.Seller, error) {
	var seller domain.Seller
	err := r.db.Get(&seller, "SELECT * FROM seller WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error get seller by id: %v", err)
	}
	return &seller, nil
}

func (r *sellerRepository) GetByOwnerID(ownerID string) (*[]domain.Seller, error) {
	sellers := make([]domain.Seller, 0)
	err := r.db.Select(&sellers, "SELECT * FROM seller WHERE owner_id = $1", ownerID)
	if err != nil {
		return nil, fmt.Errorf("error get seller by owner id: %v", err)
	}
	return &sellers, nil
}

func (r *sellerRepository) GetDashboard(sellerID string) (*domain.SellerDashboard, error) {
	var dashboard domain.SellerDashboard
	err := r.db.Get(&dashboard, `
		SELECT 
			COUNT(DISTINCT o.id) AS total_order,
			COUNT(DISTINCT u.id) AS total_order_user,
			SUM(o.product_price + o.shipping_price) AS total_order_amount,
			SUM(op.quantity) AS total_order_product_unit
		FROM 
			"order" o
		INNER JOIN 
			users u ON o.user_id = u.id
		INNER JOIN 
			order_product op ON o.id = op.order_id
		WHERE 
			o.id IN (
					SELECT 
							op_sub.order_id
					FROM 
							order_product op_sub
					INNER JOIN 
							product_size ps ON op_sub.product_size_id = ps.id
					INNER JOIN 
							product_option po ON ps.product_option_id = po.id
					INNER JOIN 
							product p ON po.product_id = p.id
					WHERE 
							p.seller_id = $1
					GROUP BY 
							op_sub.order_id
			);`, sellerID)
	if err != nil {
		return nil, fmt.Errorf("error get seller dashboard: %v", err)
	}

	return &dashboard, nil
}

func (r *sellerRepository) GetDashboardProductBySellerId(sellerID string) (*[]domain.ProductDashboard, error) {
	dashboards := make([]domain.ProductDashboard, 0)
	err := r.db.Select(&dashboards, `
		SELECT 
			p.name AS product_name,
			SUM(op.quantity) AS quantity
		FROM
			order_product op
		JOIN 
			product_size ps ON op.product_size_id = ps.id
		join 
			product_option po on po.id = ps.product_option_id 
		JOIN 
			product p ON po.product_id = p.id
		WHERE 
			p.seller_id = $1
		GROUP BY 
			p.id, p."name"
		ORDER BY 
			quantity DESC;
	`, sellerID)
	if err != nil {
		return nil, fmt.Errorf("error get seller dashboard product: %v", err)
	}

	return &dashboards, nil
}

func (r *sellerRepository) GetDashboardSizeBySellerId(sellerID string) (*[]domain.OrderSizeDashboard, error) {
	dashboards := make([]domain.OrderSizeDashboard, 0)
	err := r.db.Select(&dashboards, `
		SELECT 
			s."size" AS size,
			SUM(op.quantity) AS quantity
		FROM 
			order_product op 
		JOIN 
			product_size ps ON op.product_size_id = ps.id 
		JOIN 
			size s ON ps.size_id = s.id 
		JOIN
			product_option po ON po.id = ps.product_option_id 
		JOIN
			product p ON p.id = po.product_id 
		WHERE
			p.seller_id = $1
		GROUP by
			s."size" 
		ORDER BY 
			quantity DESC;
	`, sellerID)
	if err != nil {
		return nil, fmt.Errorf("error get seller dashboard product: %v", err)
	}

	return &dashboards, nil
}

func (r *sellerRepository) Create(seller *domain.Seller) error {
	_, err := r.db.NamedExec(`
		INSERT INTO seller 
			(id, name, description, logo_url, location, bank_account_name, bank_account_number, bank_account_provider, phone, owner_id, is_verify)
		VALUES
			(:id, :name, :description, :logo_url, :location, :bank_account_name, :bank_account_number, :bank_account_provider, :phone, :owner_id, :is_verify)`, seller)
	if err != nil {
		return fmt.Errorf("error create seller: %v", err)
	}
	return nil
}

func (r *sellerRepository) Update(seller *domain.Seller) error {
	_, err := r.db.NamedExec(`
		UPDATE
			seller
		SET 
			name = :name,
			description = :description,
			logo_url = :logo_url,
			location = :location,
			bank_account_name = :bank_account_name,
			bank_account_number = :bank_account_number,
			bank_account_provider = :bank_account_provider,
			phone = :phone,
			owner_id = :owner_id,
			is_verify = :is_verify,
			updated_at = NOW()
		WHERE
			id = :id`, seller)
	if err != nil {
		return fmt.Errorf("error update seller: %v", err)
	}
	return nil
}

func (r *sellerRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM seller WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error delete seller: %v", err)
	}
	return nil
}
