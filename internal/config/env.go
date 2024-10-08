package config

import (
	"errors"
	"log"
	"os"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func GetEnv() (*domain.ConfigEnv, error) {
	err := godotenv.Load()

	if err != nil {
		log.Default().Println(err)
		return nil, errors.New("error loading .env file")
	}

	BACKEND_PORT := os.Getenv("BACKEND_PORT")
	DATABASE_URI := os.Getenv("DATABASE_URI")

	config := &domain.ConfigEnv{
		BACKEND_PORT: BACKEND_PORT,
		DATABASE_URI: DATABASE_URI,
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	return config, nil
}
