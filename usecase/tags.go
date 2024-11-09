package usecase

import "github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"

type tagsUsecase struct {
	tagsRepo domain.TagsRepository
}

func NewTagsUsecase(tagsRepo domain.TagsRepository) domain.TagsUsecase {
	return &tagsUsecase{
		tagsRepo: tagsRepo,
	}
}

func (u *tagsUsecase) GetAll() (*[]domain.Tag, error) {
	return u.tagsRepo.GetAll()
}

func (u *tagsUsecase) GetByID(id int) (*domain.Tag, error) {
	return u.tagsRepo.GetByID(id)
}

func (u *tagsUsecase) Create(tag *domain.Tag) error {
	return u.tagsRepo.Create(tag)
}

func (u *tagsUsecase) Update(tag *domain.Tag) error {
	return u.tagsRepo.Update(tag)
}

func (u *tagsUsecase) Delete(id int) error {
	//
	return u.tagsRepo.Delete(id)
}
