package repository

import (
	"database/sql"
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type favoriteRepository struct {
	db *sqlx.DB
}

func NewFavoriteRepository(db *sqlx.DB) domain.FavoriteRepository {
	return &favoriteRepository{
		db,
	}
}

func (r *favoriteRepository) GetAll() (*[]domain.Favorite, error) {
	favorites := make([]domain.Favorite, 0)
	err := r.db.Select(&favorites, "SELECT * FROM favorite")
	if err != nil {
		return nil, fmt.Errorf("failed to get all favorites: %v", err)
	}
	return &favorites, nil
}

func (r *favoriteRepository) GetByID(id string) (*domain.Favorite, error) {
	favorite := domain.Favorite{}
	err := r.db.Get(&favorite, "SELECT * FROM favorite WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get favorite by id: %v", err)
	}
	return &favorite, nil
}

func (r *favoriteRepository) GetByUserID(id string) (*[]domain.Favorite, error) {
	favorites := make([]domain.Favorite, 0)
	err := r.db.Select(&favorites, "SELECT * FROM favorite WHERE user_id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get favorite by user id: %v", err)
	}
	return &favorites, nil
}

func (r *favoriteRepository) GetByProductID(id string) (*[]domain.Favorite, error) {
	favorites := make([]domain.Favorite, 0)
	err := r.db.Select(&favorites, "SELECT * FROM favorite WHERE product_id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get favorite by product id: %v", err)
	}
	return &favorites, nil
}

func (r *favoriteRepository) Create(favorite *domain.Favorite) error {
	_, err := r.db.NamedExec("INSERT INTO favorite (user_id, product_id) VALUES (:user_id, :product_id)", favorite)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return nil
		}
		return fmt.Errorf("failed to create favorite: %v", err)
	}
	return nil
}

func (r *favoriteRepository) Delete(user_id string, product_id string) error {
	_, err := r.db.Exec("DELETE FROM favorite WHERE user_id = $1 AND product_id = $2", user_id, product_id)
	if err != nil {
		return fmt.Errorf("failed to delete favorite: %v", err)
	}
	return nil
}
