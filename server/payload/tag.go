package payload

type CreateTagGroupDTO struct {
	LABEL string   `json:"label" validate:"required"`
	SHOW  *bool    `json:"show" validate:"required"`
	TAG   []CreateTagDTO `json:"tag" validate:"omitempty"`
}

type CreateTagDTO struct {
	LABEL        string `json:"label" validate:"required"`
}

type UpdateTagGroupDTO struct {
	LABEL string   `json:"label" validate:"required"`
	SHOW  *bool    `json:"show" validate:"required"`
	TAG   []TagDTO `json:"tag" validate:"omitempty"`
}

type TagDTO struct {
	TAG_GROUP_ID int    `json:"tagGroupId" validate:"required"`
	LABEL        string `json:"label" validate:"required"`
}