package payload

type UpdateInfoUserDTO struct {
	FIRST_NAME string `json:"firstName" validate:"required"`
	LAST_NAME  string `json:"lastName" validate:"required"`
	PHONE      string `json:"phone" validate:"required"`
}
type UpdateImageUserDTO struct {
	PROFILE_URL string `json:"profileUrl" validate:"required"`
}
