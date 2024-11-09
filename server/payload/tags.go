package payload

type CreateTagDTO struct {
	LABEL string `json:"label" validate:"required"`
	SHOW  *bool  `json:"show" validate:"required"`
}

type UpdateTagDTO struct {
	ID int `json:"id" validate:"required"`
	CreateTagDTO
}
