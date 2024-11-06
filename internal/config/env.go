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
	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET := os.Getenv("GOOGLE_CLIENT_SECRET")
	GOOGLE_REDIRECT := os.Getenv("GOOGLE_REDIRECT")
	FRONTEND_URL := os.Getenv("FRONTEND_URL")

	config := &domain.ConfigEnv{
		BACKEND_PORT:         BACKEND_PORT,
		DATABASE_URI:         DATABASE_URI,
		GOOGLE_CLIENT_ID:     GOOGLE_CLIENT_ID,
		GOOGLE_CLIENT_SECRET: GOOGLE_CLIENT_SECRET,
		GOOGLE_REDIRECT:      GOOGLE_REDIRECT,
		FRONTEND_URL:         FRONTEND_URL,
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	return config, nil
}
