package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/db"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/config"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/repository"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/usecase"
	"github.com/jmoiron/sqlx"
)

func main() {
	configEnv, err := config.GetEnv()
	if err != nil {
		log.Fatal(err)
	}

	postgres, err := db.NewConnection(configEnv.DATABASE_URI)
	if err != nil {
		log.Fatal(err)
	}

	repository := initRepository(postgres)
	usecase := initUseCase(configEnv, repository)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	fiber := server.NewFiberServer(configEnv, usecase, repository)
	go fiber.Start()

	<-signals
	fmt.Println("Server is shutting down")
	if err := fiber.Close(); err != nil {
		log.Fatal("Server is not shutting down", err)
	}
	fmt.Println("fiber was successful")

	if err := postgres.Close(); err != nil {
		log.Fatal("MySQL is not shutting down", err)
	}
	fmt.Println("db was successful")
	fmt.Println("Server was successful shutdown")
}

func initRepository(
	db *sqlx.DB,
) *domain.Repository {
	return &domain.Repository{
		UserRepository:    repository.NewUserRepository(db),
		SessionRepository: repository.NewSessionRepository(db),
	}
}

func initUseCase(config *domain.ConfigEnv, repo *domain.Repository) *domain.Usecase {
	userUsecase := usecase.NewUserUsecase(repo.UserRepository)
	sessionUsecase := usecase.NewSessionUsecase(repo.SessionRepository)
	googleUsecase := usecase.NewGoogleUsecase(config)
	authUsecase := usecase.NewAuthUsecase(googleUsecase, userUsecase, sessionUsecase)
	return &domain.Usecase{
		AuthUsecase:   authUsecase,
		GoogleUsecase: googleUsecase,
		UserUsecase:   userUsecase,
	}
}
