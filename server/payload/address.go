package payload

type CreateAddressDTO struct {
	FIRST_NAME     string `json:"firstName" db:"first_name"`
	LAST_NAME      string `json:"lastName" db:"last_name"`
	COMPANY        string `json:"company" db:"company"`
	STREET_ADDRESS string `json:"streetAddress" db:"street_address"`
	STATE          string `json:"state" db:"state"`
	COUNTRY        string `json:"country" db:"country"`
	ZIP_CODE       string `json:"zipCode" db:"zip_code"`
	EMAIL          string `json:"email" db:"email"`
	PHONE          string `json:"phone" db:"phone"`
	TYPE           string `json:"type" db:"type"`
}

type UpdateAddressDTO struct {
	ID int `json:"id" db:"id"`
	CreateAddressDTO
}
