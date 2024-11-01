package server

import (
	"log"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/validator"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/controller"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/server/middleware"
	"github.com/gofiber/contrib/swagger"
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

	s.app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "/docs/swagger",
	}))

	app := s.app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	user := app.Group("/user")
	userController := controller.NewUserController(validator, s.usecase.UserUsecase)
	user.Patch("/", middlewareAuth, userController.Update)

	auth := app.Group("/auth")
	authController := controller.NewAuthController(validator, s.config, s.usecase.AuthUsecase, s.usecase.GoogleUsecase, s.usecase.UserUsecase)
	auth.Get("/me", middlewareAuth, authController.Me)
	auth.Get("/google", authController.GetUrl)
	auth.Get("/google/callback", authController.SignInWithGoogle)
	auth.Get("/logout", middlewareAuth, authController.Logout)

	seller := app.Group("/seller")
	sellerController := controller.NewSellerController(validator, s.usecase.SellerUsecase)
	seller.Get("/", sellerController.GetAll)
	seller.Get("/owner", middlewareAuth, sellerController.GetByOwner)
	seller.Get("/:id", sellerController.GetByID)
	seller.Post("/", middlewareAuth, sellerController.Create)
	seller.Patch("/", middlewareAuth, sellerController.Update)

	product := app.Group("/product")
	productController := controller.NewProductController(validator, s.usecase.ProductUsecase, s.usecase.SellerUsecase)
	product.Get("/", productController.GetAllProductWithOptionsAndSizes)
	product.Get("/seller/:id", productController.GetBySellerID)
	product.Get("/:id", productController.GetByID)
}
