package payload

type CreateSellerDTO struct {
	NAME                  string `json:"name" validate:"required"`
	DESCRIPTION           string `json:"description" validate:"required"`
	LOGO_URL              string `json:"logoUrl" validate:"required"`
	LOCATION              string `json:"location" validate:"required"`
	BANK_ACCOUNT_NAME     string `json:"bankAccountName" validate:"required"`
	BANK_ACCOUNT_NUMBER   string `json:"bankAccountNumber" validate:"required"`
	BANK_ACCOUNT_PROVIDER string `json:"bankAccountProvider" validate:"required"`
}

type UpdateSellerDTO struct {
	ID string `json:"id" validate:"required"`
	CreateSellerDTO
}
