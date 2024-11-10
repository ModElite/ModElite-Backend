package payload

type CreateAddressDTO struct {
	FIRST_NAME     string `json:"firstName" validate:"required"`
	LAST_NAME      string `json:"lastName" validate:"required"`
	COMPANY        string `json:"company"`
	STREET_ADDRESS string `json:"streetAddress" validate:"required"`
	STATE          string `json:"state" validate:"required"`
	COUNTRY        string `json:"country" validate:"required"`
	ZIP_CODE       string `json:"zipCode" validate:"required"`
	EMAIL          string `json:"email" validate:"required"`
	PHONE          string `json:"phone" validate:"required"`
	TYPE           string `json:"type" validate:"required"`
}

type UpdateAddressDTO struct {
	ID int `json:"id" validate:"required"`
	CreateAddressDTO
}
