package domain

type ConfigEnv struct {
	BACKEND_PORT string `validate:"required"`
	DATABASE_URI string `validate:"required"`
}