package repository

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type tagsRepository struct {
	db *sqlx.DB
}

func NewTagsRepository(db *sqlx.DB) domain.TagsRepository {
	return &tagsRepository{
		db: db,
	}
}

func (tr *tagsRepository) GetAll() (*[]domain.Tag, error) {
	tags := make([]domain.Tag, 0)
	err := tr.db.Select(&tags, "SELECT * FROM tags ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("error get all tags: %w", err)
	}
	return &tags, nil
}

func (tr *tagsRepository) GetByID(id int) (*domain.Tag, error) {
	var tag domain.Tag
	err := tr.db.Get(&tag, "SELECT * FROM tags WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error get tag by id: %w", err)
	}
	return &tag, nil
}

func (tr *tagsRepository) Create(tag *domain.Tag) error {
	_, err := tr.db.NamedExec("INSERT INTO tags (label, show) VALUES (:label, :show)", tag)
	if err != nil {
		return fmt.Errorf("error create tag: %w", err)
	}
	return nil
}

func (tr *tagsRepository) Update(tag *domain.Tag) error {
	_, err := tr.db.NamedExec("UPDATE tags SET label = :label, show = :show WHERE id = :id", tag)
	if err != nil {
		return fmt.Errorf("error update tag: %w", err)
	}
	return nil
}

func (tr *tagsRepository) Delete(id int) error {
	_, err := tr.db.Exec("DELETE FROM tags WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error delete tag: %w", err)
	}
	return nil
}
