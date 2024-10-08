package controller

import (
	"fmt"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	config        *domain.ConfigEnv
	authUsecase   domain.AuthUsecase
	googleUsecase domain.GoogleUsecase
}

func NewAuthController(config *domain.ConfigEnv, authUsecase domain.AuthUsecase, googleUsecase domain.GoogleUsecase) *authController {
	return &authController{
		config:        config,
		authUsecase:   authUsecase,
		googleUsecase: googleUsecase,
	}
}

func (auth *authController) GetUrl(c *fiber.Ctx) error {
	path := auth.googleUsecase.GoogleConfig()
	url := path.AuthCodeURL("state")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"url": url})
}

func (auth *authController) SignInWithGoogle(ctx *fiber.Ctx) error {
	cookie, err := auth.authUsecase.SignInWithGoogle(ctx)
	if err != nil {
		fmt.Println("SignIn :", err)
		return ctx.Status(500).Send([]byte("Internal Server Error"))
	}

	ctx.Cookie(cookie)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Success"})
}

func (auth *authController) SignOut(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Success"})
}
