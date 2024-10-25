package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	config        *domain.ConfigEnv
	authUsecase   domain.AuthUsecase
	googleUsecase domain.GoogleUsecase
	userUsecase   domain.UserUsecase
}

func NewAuthController(
	config *domain.ConfigEnv,
	authUsecase domain.AuthUsecase,
	googleUsecase domain.GoogleUsecase,
	userUsecase domain.UserUsecase,
) *authController {
	return &authController{
		config:        config,
		authUsecase:   authUsecase,
		googleUsecase: googleUsecase,
		userUsecase:   userUsecase,
	}
}

func (auth *authController) Me(ctx *fiber.Ctx) error {
	user, err := auth.userUsecase.Get(ctx.Locals(constant.USER_ID).(string))
	if err != nil {
		return ctx.Status(500).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	} else if user == nil {
		return ctx.Status(403).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_UNAUTHORIZED,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    user,
	})

}

func (auth *authController) GetUrl(c *fiber.Ctx) error {
	path := auth.googleUsecase.GoogleConfig()
	url := path.AuthCodeURL("state")

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    url,
	})
}

func (auth *authController) SignInWithGoogle(ctx *fiber.Ctx) error {
	cookie, err := auth.authUsecase.SignInWithGoogle(ctx)
	if err != nil {
		return ctx.Status(500).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})

	}

	ctx.Cookie(cookie)
	return ctx.Redirect("http://localhost:3000")
}

func (auth *authController) SignOut(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}
