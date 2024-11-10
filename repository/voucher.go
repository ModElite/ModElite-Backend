package repository

import (
	"database/sql"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/google/uuid"
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

func (r *voucherRepository) Search(code string) (*domain.Voucher, error) {
	voucher := domain.Voucher{}
	err := r.db.Get(&voucher, `
	SELECT
		* 
	FROM
		voucher 
	WHERE
		( SELECT COUNT (id) FROM "order" WHERE "order".voucher_code = voucher."id" ) < voucher.quota 
		AND expired_at > NOW( ) AND "voucher".code = LOWER($1) LIMIT 1;
	`, code)

	if err != nil {
		return nil, err
	}

	return &voucher, nil
}

func (r *voucherRepository) CheckDuplicateCode(code string) bool {
	voucher := domain.Voucher{}
	err := r.db.Get(&voucher, `
	SELECT
		* 
	FROM
		voucher 
	WHERE
		code = $1 AND expired_at > NOW( ) LIMIT 1;
	`, code)

	return err != sql.ErrNoRows
}

func (r *voucherRepository) CreateVoucher(voucher *domain.Voucher) error {
	voucher.ID = uuid.New().String()

	_, err := r.db.NamedExec(`
	INSERT INTO voucher (id,code, min_total_price, max_discount, percentage, quota, expired_at) 
	VALUES (:id, LOWER(:code), :min_total_price, :max_discount, :percentage, :quota, :expired_at);
	`, voucher)

	if err != nil {
		return err
	}

	return nil
}
