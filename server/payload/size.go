package payload

type SizeDTO struct {
	SIZE string `json:"size" validate:"required"`
}
