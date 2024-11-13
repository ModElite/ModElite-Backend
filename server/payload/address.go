package payload

type AddressDTO struct {
	FIRST_NAME   string `json:"firstName" validate:"required"`
	LAST_NAME    string `json:"lastName" validate:"required"`
	EMAIL        string `json:"email" validate:"required"`
	PHONE        string `json:"phone" validate:"required"`
	LABEL        string `json:"label" validate:"required"`
	DEFAULT      *bool  `json:"default" validate:"required"`
	ADDRESS      string `json:"address" validate:"required"`
	SUB_DISTRICT string `json:"subDistrict" validate:"required"`
	DISTRICT     string `json:"district" validate:"required"`
	PROVINCE     string `json:"province" validate:"required"`
	ZIP_CODE     string `json:"zipCode" validate:"required"`
}
