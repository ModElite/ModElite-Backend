package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *domain.User) error {
	_, err := r.db.NamedExec(
		`INSERT INTO users (id, email, google_id, first_name, last_name, phone, profile_url, created_at, updated_at)`+
			`VALUES (:id, :email, :google_id, :first_name, :last_name, :phone, :profile_url, :created_at, :updated_at)`,
		user,
	)
	if err != nil {
		return fmt.Errorf("cannot query to create user: %w", err)
	}
	return nil
}

func (r *userRepository) Get(id string) (*domain.User, error) {
	user := domain.User{}
	err := r.db.Get(&user, `SELECT * FROM users WHERE id = $1`, id)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get user by id: %w", err)
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	user := domain.User{}
	err := r.db.Get(&user, `SELECT * FROM users WHERE email = $1`, email)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get user by email: %w", err)
	}
	return &user, nil
}

func (r *userRepository) Update(user *domain.User) error {
	user.UPDATED_AT = time.Now()
	_, err := r.db.NamedExec(
		`UPDATE users SET email = :email, google_id = :google_id, first_name = :first_name, last_name = :last_name, phone = :phone, profile_url = :profile_url, updated_at = :updated_at WHERE id = :id`,
		user,
	)
	if err != nil {
		return fmt.Errorf("cannot query to update user: %w", err)
	}
	return nil
}
