package usecase

import "github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"

type tagsUsecase struct {
	tagsRepo        domain.TagsRepository
	productTagsRepo domain.ProductTagsRepository
}

func NewTagsUsecase(
	tagsRepo domain.TagsRepository,
	productTagsRepo domain.ProductTagsRepository,
) domain.TagsUsecase {
	return &tagsUsecase{
		tagsRepo:        tagsRepo,
		productTagsRepo: productTagsRepo,
	}
}

func (u *tagsUsecase) GetAll() (*[]domain.Tag, error) {
	return u.tagsRepo.GetAll()
}

func (u *tagsUsecase) GetByID(id int) (*domain.Tag, error) {
	return u.tagsRepo.GetByID(id)
}

func (u *tagsUsecase) GetByProductID(productId string) (*[]domain.Tag, error) {
	productTags, err := u.productTagsRepo.GetByProductID(productId)
	if err != nil {
		return nil, err
	}

	tags := make([]domain.Tag, 0)
	for _, productTag := range *productTags {
		tag, err := u.tagsRepo.GetByID(productTag.TAG_ID)
		if err != nil {
			return nil, err
		}
		tags = append(tags, *tag)
	}

	return &tags, nil
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
