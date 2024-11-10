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
		AllowOrigins:     "http://localhost:5173,http://localhost:3000,https://test.sssboom.xyz",
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
	middlewareAuthAdmin := middleware.NewAuthAdminMiddleware(s.usecase.UserUsecase)

	s.app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "/docs/swagger",
	}))

	app := s.app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	address := app.Group("/address")
	addressController := controller.NewAddressController(validator, s.usecase.AddressUsecase, s.usecase.UserUsecase)
	address.Get("/", middlewareAuth, addressController.GetByUserID)
	address.Post("/", middlewareAuth, addressController.Create)
	address.Put("/", middlewareAuth, addressController.Update)
	address.Delete("/:id", middlewareAuth, addressController.Delete)

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
	productController := controller.NewProductController(validator, s.usecase.ProductUsecase, s.usecase.SellerUsecase, s.usecase.UserUsecase)
	product.Get("/", productController.GetAllProductWithOptionsAndSizes)
	product.Get("/seller/:id", productController.GetBySellerID)
	product.Get("/:id", productController.GetByID)
	product.Post("/", middlewareAuth, productController.Create)

	favorite := app.Group("/favorite")
	favoriteController := controller.NewFavoriteController(validator, s.usecase.FavoriteUsecase)
	favorite.Get("/", middlewareAuth, favoriteController.GetByUserID)
	favorite.Post("/", middlewareAuth, favoriteController.Create)
	favorite.Delete("/:id", middlewareAuth, favoriteController.Delete)

	tags := app.Group("/tags")
	tagsController := controller.NewTagsController(validator, s.usecase.TagsUsecase)
	tags.Get("/", tagsController.GetTags)
	tags.Post("/", middlewareAuth, middlewareAuthAdmin, tagsController.CreateTag)
	tags.Patch("/", middlewareAuth, middlewareAuthAdmin, tagsController.UpdateTag)

	order := app.Group("/order")
	orderController := controller.NewOrderController(validator, s.usecase.OrderUsecase, s.usecase.VoucherUsecase)
	order.Get("/", orderController.GetAll)
	order.Get("/self", middlewareAuth, orderController.GetSelfOrder)
	order.Post("/", middlewareAuth, orderController.CreateOrder)
}
