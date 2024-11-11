package payload

type TagGroupDTO struct {
	LABEL string   `json:"label" validate:"required"`
	SHOW  *bool    `json:"show" validate:"required"`
	TAG   []TagDTO `json:"tag" validate:"omitempty"`
}

type TagDTO struct {
	TAG_GROUP_ID int    `json:"tagGroupId" validate:"required"`
	LABEL        string `json:"label" validate:"required"`
}
