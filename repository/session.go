package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type sessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) domain.SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (r *sessionRepository) Create(session *domain.Session) error {
	_, err := r.db.NamedExec(`INSERT INTO session (id, user_id, ip_address, user_agent, expired_at, created_at)`+
		`VALUES (:id, :user_id, :ip_address, :user_agent, :expired_at, :created_at)`, session)
	if err != nil {
		return fmt.Errorf("error creating session: %w", err)
	}
	return nil
}

func (r *sessionRepository) GetByID(id string) (*domain.Session, error) {
	session := domain.Session{}
	err := r.db.Get(&session, `SELECT * FROM session WHERE id = $1 LIMIT 1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error getting session: %w", err)
	}
	return &session, nil
}

func (r *sessionRepository) ExtendExpiredAt(id string) error {
	_, err := r.db.Exec(`UPDATE session SET expired_at = $1 WHERE id = $2`, time.Now().Add(time.Hour*24*3), id)
	if err != nil {
		return fmt.Errorf("error extending session: %w", err)
	}
	return nil
}

func (r *sessionRepository) DeleteById(id string) error {
	_, err := r.db.Exec(`DELETE FROM session WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting session: %w", err)
	}
	return nil
}
