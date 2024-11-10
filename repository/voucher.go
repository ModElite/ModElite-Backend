package repository

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type voucherRepository struct {
	db *sqlx.DB
}

func NewVoucherRepository(db *sqlx.DB) domain.VoucherRepository {
	return &voucherRepository{
		db: db,
	}
}

func (r *voucherRepository) GetByID(id string) (*domain.Voucher, error) {
	voucher := domain.Voucher{}
	err := r.db.Get(&voucher, `
	SELECT
		* 
	FROM
		voucher 
	WHERE
		( SELECT COUNT (id) FROM "order" WHERE "order".voucher_code = voucher."id" ) < voucher.quota 
		AND expired_at > NOW( ) AND "voucher".id = $1 LIMIT 1;
	`, id)
	// if voucher is not found

	if err != nil {
		return nil, err
	}

	return &voucher, nil
}
