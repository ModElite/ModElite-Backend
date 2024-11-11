package repository

import (
	"database/sql"
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type tagGroupRepository struct {
	db *sqlx.DB
}

func NewTagGroupRepository(db *sqlx.DB) domain.TagGroupRepository {
	return &tagGroupRepository{
		db: db,
	}
}

func (r *tagGroupRepository) GetAll() (*[]domain.TagGroup, error) {
	tagGroups := make([]domain.TagGroup, 0)
	err := r.db.Select(&tagGroups, "SELECT * FROM tag_group ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("error get all tag group: %w", err)
	}
	return &tagGroups, nil
}

func (r *tagGroupRepository) GetByID(id int) (*domain.TagGroup, error) {
	tagGroup := domain.TagGroup{}
	err := r.db.Get(&tagGroup, "SELECT * FROM tag_group WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error get tag group by id: %w", err)
	}
	return &tagGroup, nil
}

func (r *tagGroupRepository) Create(tagGroup *domain.TagGroup) (*int, error) {
	tx := r.db.MustBegin()
	var lastInsertID int
	query := `INSERT INTO tag_group (label, show) VALUES ($1, $2) RETURNING id`
	err := tx.QueryRowx(query, tagGroup.LABEL, tagGroup.SHOW).Scan(&lastInsertID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, fmt.Errorf("error rollback create tag group: %w", err)
		}
		return nil, fmt.Errorf("error create tag group: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error commit create tag group: %w", err)
	}

	return &lastInsertID, nil
}

func (r *tagGroupRepository) Update(tagGroup *domain.TagGroup) error {
	_, err := r.db.NamedExec("UPDATE tag_group SET label = :label, show = :show WHERE id = :id", tagGroup)
	if err != nil {
		return fmt.Errorf("error update tag group: %w", err)
	}
	return nil
}

func (r *tagGroupRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM tag_group WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error delete tag group: %w", err)
	}
	return nil
}
