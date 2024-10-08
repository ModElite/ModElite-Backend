package server

import (
	"log"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/controller"
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
		AllowOrigins:     "http://localhost:5173",
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
	app := s.app

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	authController := controller.NewAuthController(s.config, s.usecase.AuthUsecase, s.usecase.GoogleUsecase)
	app.Get("/auth/google/url", authController.GetUrl)
}
