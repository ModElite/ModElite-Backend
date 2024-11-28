package repository

import (
	"database/sql"
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type tagRepository struct {
	db *sqlx.DB
}

func NewTagRepository(db *sqlx.DB) domain.TagRepository {
	return &tagRepository{
		db: db,
	}
}

func (r *tagRepository) GetAll() (*[]domain.Tag, error) {
	tags := make([]domain.Tag, 0)
	err := r.db.Select(&tags, "SELECT * FROM tag ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("error get all tags: %w", err)
	}
	return &tags, nil
}

func (r *tagRepository) GetByID(id int) (*domain.Tag, error) {
	var tag domain.Tag
	err := r.db.Get(&tag, "SELECT * FROM tag WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error get tag by id: %w", err)
	}
	return &tag, nil
}

func (r *tagRepository) GetByTagGroupID(tagGroupID int) (*[]domain.Tag, error) {
	tags := make([]domain.Tag, 0)
	err := r.db.Select(&tags, "SELECT * FROM tag WHERE tag_group_id = $1 ORDER BY id", tagGroupID)
	if err != nil {
		return nil, fmt.Errorf("error get tags by tag group id: %w", err)
	}
	return &tags, nil
}

func (r *tagRepository) Create(tag *domain.Tag) (*int, error) {
	tx := r.db.MustBegin()
	var lastInsertID int
	query := `INSERT INTO tag (label, tag_group_id, image_url) VALUES ($1, $2, $3) RETURNING id`
	err := tx.QueryRowx(query, tag.LABEL, tag.TAG_GRUOP_ID, tag.IMAGE_URL).Scan(&lastInsertID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, fmt.Errorf("error rollback create tag: %w", err)
		}
		return nil, fmt.Errorf("error create tag: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error commit create tag: %w", err)
	}

	return &lastInsertID, nil
}

func (r *tagRepository) Update(tag *domain.Tag) error {
	_, err := r.db.NamedExec("UPDATE tag SET label = :label, tag_group_id = :tag_group_id, image_url = :image_url WHERE id = :id", tag)
	if err != nil {
		return fmt.Errorf("error update tag: %w", err)
	}
	return nil
}

func (r *tagRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM tag WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error delete tag: %w", err)
	}
	return nil
}
