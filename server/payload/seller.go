package payload

type CreateSellerDTO struct {
	NAME        string `json:"name" validate:"required"`
	DESCRIPTION string `json:"description" validate:"required"`
	LOGO_URL    string `json:"logoUrl" validate:"required"`
	LOCATION    string `json:"location" validate:"required"`
}

type UpdateSellerDTO struct {
	ID string `json:"id" validate:"required"`
	CreateSellerDTO
}
