package server

import (
	"log"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/validator"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/controller"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	app        *fiber.App
	config     *domain.ConfigEnv
	usecase    *domain.Usecase
	repository *domain.Repository
}

func NewFiberServer(
	config *domain.ConfigEnv,
	usecase *domain.Usecase,
	repository *domain.Repository,
) *FiberServer {
	return &FiberServer{
		config:     config,
		usecase:    usecase,
		repository: repository,
	}
}

func (s *FiberServer) Start() {
	app := fiber.New(fiber.Config{
		AppName: "API",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173,http://localhost:3000",
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
	}))

	s.app = app
	s.Route()

	if err := app.Listen(":" + string(s.config.BACKEND_PORT)); err != nil {
		log.Fatal("Server is not running")
	}
}

func (s *FiberServer) Close() error {
	return s.app.Shutdown()
}

func (s *FiberServer) Route() {
	validator := validator.NewPayloadValidator()
	middlewareAuth := middleware.NewAuthMiddleware(s.usecase.SessionUsecase)

	app := s.app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	auth := app.Group("/auth")
	authController := controller.NewAuthController(validator, s.config, s.usecase.AuthUsecase, s.usecase.GoogleUsecase, s.usecase.UserUsecase)
	auth.Get("/me", middlewareAuth, authController.Me)
	auth.Get("/google", authController.GetUrl)
	auth.Get("/google/callback", authController.SignInWithGoogle)
	auth.Get("/logout", middlewareAuth, authController.Logout)

	seller := app.Group("/seller")
	sellerController := controller.NewSellerController(validator, s.usecase.SellerUsecase)
	seller.Get("/", middlewareAuth, sellerController.GetAll)
}
