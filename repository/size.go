package repository

import (
	"fmt"
	"time"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type sizeRepository struct {
	db *sqlx.DB
}

func NewSizeRepository(db *sqlx.DB) domain.SizeRepository {
	return &sizeRepository{
		db: db,
	}
}

func (r *sizeRepository) GetAll() (*[]domain.Size, error) {
	sizes := make([]domain.Size, 0)
	err := r.db.Select(&sizes, "SELECT * FROM size")
	if err != nil {
		return nil, fmt.Errorf("error cannot get sizes: %w", err)
	}
	return &sizes, nil
}

func (r *sizeRepository) GetByID(id string) (*domain.Size, error) {
	var size domain.Size
	err := r.db.Get(&size, "SELECT * FROM size WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error cannot get size: %w", err)
	}
	return &size, nil
}

func (r *sizeRepository) Create(size *domain.Size) error {
	_, err := r.db.NamedExec("INSERT INTO size (id, size) VALUES (:id, :size)", size)
	if err != nil {
		return fmt.Errorf("error cannot create size: %w", err)
	}
	return nil
}

func (r *sizeRepository) Update(size *domain.Size) error {
	size.UPDATED_AT = time.Now()
	_, err := r.db.NamedExec("UPDATE size SET size = :size, updated_at = :updated_at WHERE id = :id", size)
	if err != nil {
		return fmt.Errorf("error cannot update size: %w", err)
	}
	return nil
}

func (r *sizeRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM size WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error cannot delete size: %w", err)
	}
	return nil
}
