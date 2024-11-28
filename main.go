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

//	@title			Sofware Engineering Project Backend API
//	@version		1.0
//	@description	This is a sample server celler server.

// @BasePath	/
// @securityDefinitions.apikey 	ApiKeyAuth
// @in 													header
// @name												Authorization
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
		AddressRepository:           repository.NewAddressRepository(db),
		UserRepository:              repository.NewUserRepository(db),
		SessionRepository:           repository.NewSessionRepository(db),
		SellerRepository:            repository.NewSellerRepository(db),
		SellerTransactionRepository: repository.NewSellerTransactionRepository(db),
		ProductRepository:           repository.NewProductRepository(db),
		ProductOptionRepository:     repository.NewProductOptionRepository(db),
		ProductSizeRepository:       repository.NewProductSizeRepository(db),
		SizeRepository:              repository.NewSizeRepository(db),
		FavoriteRepository:          repository.NewFavoriteRepository(db),
		TagRepository:               repository.NewTagRepository(db),
		TagGroupRepository:          repository.NewTagGroupRepository(db),
		ProductTagRepository:        repository.NewProductTagRepository(db),
		OrderRepository:             repository.NewOrderRepository(db),
		VoucherRepository:           repository.NewVoucherRepository(db),
		GeoLocationRepository:       repository.NewGeoLocationRepository(db),
		CartRepository:              repository.NewCartRepository(db),
	}
}

func initUseCase(
	config *domain.ConfigEnv,
	repo *domain.Repository,
) *domain.Usecase {
	googleUsecase := usecase.NewGoogleUsecase(config)
	userUsecase := usecase.NewUserUsecase(repo.UserRepository)
	addressUsecase := usecase.NewAddressUsecase(repo.AddressRepository, userUsecase)
	sessionUsecase := usecase.NewSessionUsecase(repo.SessionRepository)
	authUsecase := usecase.NewAuthUsecase(googleUsecase, userUsecase, sessionUsecase)
	sellerUsecase := usecase.NewSellerUsecase(repo.SellerRepository, userUsecase)
	sellerTransactionUsecase := usecase.NewSellerTransactionUsecase(repo.SellerTransactionRepository)
	productOptionUsecase := usecase.NewProductOptionUsecase(repo.ProductOptionRepository, repo.ProductSizeRepository)
	productSizeUsecase := usecase.NewProductSizeUsecase(repo.ProductSizeRepository)
	sizeUsecase := usecase.NewSizeUsecase(repo.SizeRepository)
	tagUsecase := usecase.NewTagUsecase(repo.TagRepository, repo.TagGroupRepository, repo.ProductTagRepository)
	productUsecase := usecase.NewProductUsecase(repo.ProductRepository, productOptionUsecase, productSizeUsecase, tagUsecase, userUsecase, sellerUsecase)
	favoriteUsecase := usecase.NewFavoriteUsecase(repo.FavoriteRepository, productUsecase)
	vocherUsecase := usecase.NewVoucherUsecase(repo.VoucherRepository)
	orderUsecase := usecase.NewOrderUsecase(repo.OrderRepository, repo.ProductRepository)
	geoLocationUsecase := usecase.NewGeoLocationUseCase(repo.GeoLocationRepository)
	cartUsecase := usecase.NewCartUsecase(repo.CartRepository)
	return &domain.Usecase{
		AddressUsecase:           addressUsecase,
		AuthUsecase:              authUsecase,
		GoogleUsecase:            googleUsecase,
		UserUsecase:              userUsecase,
		SessionUsecase:           sessionUsecase,
		SellerUsecase:            sellerUsecase,
		SellerTransactionUsecase: sellerTransactionUsecase,
		ProductUsecase:           productUsecase,
		ProductOptionUsecase:     productOptionUsecase,
		ProductSizeUsecase:       productSizeUsecase,
		SizeUsecase:              sizeUsecase,
		FavoriteUsecase:          favoriteUsecase,
		TagUsecase:               tagUsecase,
		OrderUsecase:             orderUsecase,
		VoucherUsecase:           vocherUsecase,
		GeoLocationUsecase:       geoLocationUsecase,
		CartUsecase:              cartUsecase,
	}
}
