package usecase

import (
	"sync"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
)

type tagUsecase struct {
	tagRepo        domain.TagRepository
	tagGroupRepo   domain.TagGroupRepository
	productTagRepo domain.ProductTagRepository
}

func NewTagUsecase(
	tagRepo domain.TagRepository,
	tagGroupRepo domain.TagGroupRepository,
	productTagRepo domain.ProductTagRepository,
) domain.TagUsecase {
	return &tagUsecase{
		tagRepo:        tagRepo,
		tagGroupRepo:   tagGroupRepo,
		productTagRepo: productTagRepo,
	}
}
func (u *tagUsecase) GetAllTagGroup() (*[]domain.TagGroup, error) {
	return u.tagGroupRepo.GetAll()
}

func (u *tagUsecase) GetAllTagGroupWithTags() (*[]domain.TagGroup, error) {
	tagGroups, err := u.tagGroupRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, len(*tagGroups))
	for i, tagGroup := range *tagGroups {
		wg.Add(1)
		go func(i int, tagGroup domain.TagGroup) {
			defer wg.Done()
			tags, err := u.tagRepo.GetByTagGroupID(tagGroup.ID)
			if err != nil {
				errChan <- err
				return
			}

			mu.Lock()
			(*tagGroups)[i].TAG = tags
			mu.Unlock()
		}(i, tagGroup)
	}
	wg.Wait()
	close(errChan)
	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return tagGroups, nil
}

func (u *tagUsecase) CreateTagGroup(TagGroup *domain.TagGroup) error {
	id, err := u.tagGroupRepo.Create(TagGroup)
	if TagGroup.TAG != nil {
		for _, item := range *TagGroup.TAG {
			if _, err = u.tagRepo.Create(&domain.Tag{
				LABEL:        item.LABEL,
				TAG_GRUOP_ID: *id,
			}); err != nil {
				return err
			}
		}
	}
	return err
}

func (u *tagUsecase) UpdateTagGroup(TagGroup *domain.TagGroup) error {
	return u.tagGroupRepo.Update(TagGroup)
}

func (u *tagUsecase) DeleteTagGroup(id int) error {
	return u.tagGroupRepo.Delete(id)
}

func (u *tagUsecase) GetAllTag(optionalParams ...func(*domain.Tag) bool) (*[]domain.Tag, error) {
	tags, err := u.tagRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(optionalParams) == 0 {
		return tags, nil
	}

	filteredTags := make([]domain.Tag, 0)
	for _, tag := range *tags {
		include := true
		for _, param := range optionalParams {
			if !param(&tag) {
				include = false
				break
			}
		}
		if include {
			filteredTags = append(filteredTags, tag)
		}
	}

	return &filteredTags, nil
}

func (u *tagUsecase) CreateTag(tag *domain.Tag) error {
	_, err := u.tagRepo.Create(tag)
	return err
}

func (u *tagUsecase) GetTagByID(id int) (*domain.Tag, error) {
	return u.tagRepo.GetByID(id)
}

func (u *tagUsecase) UpdateTag(tag *domain.Tag) error {
	return u.tagRepo.Update(tag)
}

func (u *tagUsecase) DeleteTag(id int) error {
	return u.tagRepo.Delete(id)
}

func (u *tagUsecase) GetTagByProductID(productID string) (*[]domain.Tag, error) {
	productTags, err := u.productTagRepo.GetByProductID(productID)
	if err != nil {
		return nil, err
	}

	tags := make([]domain.Tag, 0)
	for _, productTag := range *productTags {
		tag, err := u.tagRepo.GetByID(productTag.TAG_ID)
		if err != nil {
			return nil, err
		}
		tags = append(tags, *tag)
	}

	return &tags, nil
}

func (u *tagUsecase) CreateProductTag(productID string, tagID int) error {
	tag, err := u.tagRepo.GetByID(tagID)
	if err != nil || tag == nil {
		return err
	}
	return u.productTagRepo.Create(productID, tagID)
}
