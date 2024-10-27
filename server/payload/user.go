package payload

type UpdateUserDTO struct {
	FIRST_NAME  string `json:"firstName" validate:"required"`
	LAST_NAME   string `json:"lastName" validate:"required"`
	PHONE       string `json:"phone"`
	PROFILE_URL string `json:"profileUrl"`
}
