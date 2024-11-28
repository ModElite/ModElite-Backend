package domain

type TagGroup struct {
	ID    int    `json:"id" db:"id"`
	LABEL string `json:"label" db:"label"`
	SHOW  bool   `json:"show" db:"show"`
	TAG   *[]Tag `json:"tag" db:"-"`
}

type TagGroupRepository interface {
	GetAll() (*[]TagGroup, error)
	GetByID(id int) (*TagGroup, error)
	Create(tagGroup *TagGroup) (*int, error)
	Update(tagGroup *TagGroup) error
	Delete(id int) error
}

type Tag struct {
	ID           int    `json:"id" db:"id"`
	TAG_GRUOP_ID int    `json:"tagGroupId" db:"tag_group_id"`
	LABEL        string `json:"label" db:"label"`
	IMAGE_URL    string `json:"imageUrl" db:"image_url"`
}

type TagRepository interface {
	GetAll() (*[]Tag, error)
	GetByID(id int) (*Tag, error)
	GetByTagGroupID(tagGroupID int) (*[]Tag, error)
	Create(tag *Tag) (*int, error)
	Update(tag *Tag) error
	Delete(id int) error
}

type TagUsecase interface {
	GetAllTagGroup() (*[]TagGroup, error)
	GetAllTagGroupWithTags() (*[]TagGroup, error)
	CreateTagGroup(TagGroup *TagGroup) error
	UpdateTagGroup(TagGroup *TagGroup) error
	DeleteTagGroup(id int) error
	GetAllTag(optionalParams ...func(*Tag) bool) (*[]Tag, error)
	GetTagByID(id int) (*Tag, error)
	CreateTag(tag *Tag) error
	UpdateTag(tag *Tag) error
	DeleteTag(id int) error
	GetTagByProductID(productID string) (*[]Tag, error)
	CreateProductTag(productID string, tagID int) error
}
