package server

import (
	"log"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/docs"
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
		AllowOrigins:     "http://localhost:5173,http://localhost:3000,https://modelite.sssboom.xyz",
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

	if s.config.APP_ENV == "PRODUCTION" {
		docs.SwaggerInfo.Schemes = []string{"https"}
		docs.SwaggerInfo.Host = "test.sssboom.xyz"
	} else {
		docs.SwaggerInfo.Schemes = []string{"http"}
		docs.SwaggerInfo.Host = "localhost:8080"
	}

	app := s.app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	address := app.Group("/address")
	addressController := controller.NewAddressController(validator, s.usecase.AddressUsecase, s.usecase.UserUsecase)
	address.Get("/", middlewareAuth, addressController.GetByUserID)
	address.Get("/:id", middlewareAuth, addressController.GetByID)
	address.Post("/", middlewareAuth, addressController.Create)
	address.Put("/:id", middlewareAuth, addressController.Update)
	address.Delete("/:id", middlewareAuth, addressController.Delete)

	user := app.Group("/user")
	userController := controller.NewUserController(validator, s.usecase.UserUsecase)
	user.Patch("/", middlewareAuth, userController.UpdateInfo)
	user.Patch("/profile", middlewareAuth, userController.UpdateImage)

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
	productController := controller.NewProductController(validator, s.usecase.ProductUsecase, s.usecase.SellerUsecase, s.usecase.UserUsecase, s.usecase.TagUsecase)
	product.Get("/", productController.GetAllProductWithOptionsAndSizes)
	product.Get("/seller/:id", productController.GetBySellerID)
	product.Get("/:id", productController.GetByID)
	product.Post("/", middlewareAuth, productController.Create)
	product.Put("/:id", middlewareAuth, productController.Update)
	product.Delete("/:id", middlewareAuth, productController.SoftDelete)

	favorite := app.Group("/favorite")
	favoriteController := controller.NewFavoriteController(validator, s.usecase.FavoriteUsecase)
	favorite.Get("/", middlewareAuth, favoriteController.GetByUserID)
	favorite.Post("/", middlewareAuth, favoriteController.Create)
	favorite.Delete("/:id", middlewareAuth, favoriteController.Delete)

	tag := app.Group("/tag")
	tagController := controller.NewTagController(validator, s.usecase.TagUsecase)
	tag.Get("/", tagController.GetAllTag)
	tag.Get("/:id", tagController.GetTag)
	tag.Post("/", middlewareAuth, middlewareAuthAdmin, tagController.CreateTag)
	tag.Put("/:id", middlewareAuth, middlewareAuthAdmin, tagController.UpdateTag)
	tag.Delete("/:id", middlewareAuth, middlewareAuthAdmin, tagController.DeleteTag)
	tagGroup := tag.Group("/group")
	tagGroup.Get("/", tagController.GetAllTagGroup)
	tagGroup.Get("/:id", tagController.GetTag)
	tagGroup.Post("/", middlewareAuth, middlewareAuthAdmin, tagController.CreateTagGroup)
	tagGroup.Put("/:id", middlewareAuth, middlewareAuthAdmin, tagController.UpdateTagGroup)
	tagGroup.Delete("/:id", middlewareAuth, middlewareAuthAdmin, tagController.DeleteTagGroup)

	order := app.Group("/order")
	orderController := controller.NewOrderController(validator, s.usecase.OrderUsecase, s.usecase.VoucherUsecase)
	order.Get("/", orderController.GetAll)
	order.Get("/self", middlewareAuth, orderController.GetSelfOrder)
	order.Get("/self/:id", middlewareAuth, orderController.GetSelfOrderDetail)
	order.Post("/", middlewareAuth, orderController.CreateOrder)

	voucher := app.Group("/voucher")
	voucherController := controller.NewVoucherController(validator, s.usecase.VoucherUsecase)
	voucher.Get("/:code", middlewareAuth, voucherController.Search)
	voucher.Post("/", middlewareAuth, middlewareAuthAdmin, voucherController.CreateVoucher)

	geoLocation := app.Group("/geo-location")
	geoLocationController := controller.NewGeoLocationController(validator, s.usecase.GeoLocationUsecase)
	geoLocation.Get("/provinces", geoLocationController.GetProvinces)
	geoLocation.Get("/districts/:province_id", geoLocationController.GetDistrictsByProvinceId)
	geoLocation.Get("/sub-districts/:district_id", geoLocationController.GetSubDistrictsByDistrictId)

	cart := app.Group("/cart")
	cartController := controller.NewCartController(validator, s.usecase.CartUsecase)
	cart.Get("/", middlewareAuth, middlewareAuthAdmin, cartController.GetAll)
	cart.Get("/self", middlewareAuth, cartController.GetCartByUserId)
	cart.Post("/", middlewareAuth, cartController.EditCart)

	uploadController := controller.NewUploadController(s.config)
	app.Get("/upload/:filename", uploadController.GetFile)
	app.Post("/upload", uploadController.UploadFile)
}
