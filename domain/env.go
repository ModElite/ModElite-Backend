package domain

type ConfigEnv struct {
	APP_ENV              string `validate:"required"`
	BACKEND_PORT         string `validate:"required"`
	DATABASE_URI         string `validate:"required"`
	GOOGLE_CLIENT_ID     string `validate:"required"`
	GOOGLE_CLIENT_SECRET string `validate:"required"`
	GOOGLE_REDIRECT      string `validate:"required"`
	FRONTEND_URL         string `validate:"required"`
}
