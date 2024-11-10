package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	validator     domain.ValidatorUsecase
	config        *domain.ConfigEnv
	authUsecase   domain.AuthUsecase
	googleUsecase domain.GoogleUsecase
	userUsecase   domain.UserUsecase
}

func NewAuthController(
	validator domain.ValidatorUsecase,
	config *domain.ConfigEnv,
	authUsecase domain.AuthUsecase,
	googleUsecase domain.GoogleUsecase,
	userUsecase domain.UserUsecase,
) *authController {
	return &authController{
		validator:     validator,
		config:        config,
		authUsecase:   authUsecase,
		googleUsecase: googleUsecase,
		userUsecase:   userUsecase,
	}
}

// @Summary Get user profile
// @Description Get user profile
// @Tags Auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} domain.Response
// @Router /api/auth/me [get]
func (auth *authController) Me(ctx *fiber.Ctx) error {
	user, err := auth.userUsecase.Get(ctx.Locals(constant.USER_ID).(string))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	} else if user == nil {
		return ctx.Status(fiber.StatusForbidden).JSON(domain.Response{
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

// @Summary Get google auth url
// @Description Get google auth url
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response
// @Router /api/auth/google [get]
// @Param redirectUrl query string false "Redirect url"
func (auth *authController) GetUrl(ctx *fiber.Ctx) error {
	path := auth.googleUsecase.GoogleConfig()
	redirectUrl := ctx.Query("redirectUrl", auth.config.FRONTEND_URL)
	url := path.AuthCodeURL(redirectUrl)

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    url,
	})
}

// @Summary Sign in with google
// @Description Sign in with google
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response
// @Router /api/auth/google/callback [get]
func (auth *authController) SignInWithGoogle(ctx *fiber.Ctx) error {
	stateUrl := ctx.Query("state")
	cookie, err := auth.authUsecase.SignInWithGoogle(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})

	}

	ctx.Cookie(cookie)
	return ctx.Redirect(stateUrl + "?token=" + cookie.Value)
}

// @Summary Logout
// @Description Logout
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response
// @Router /api/auth/logout [get]
func (auth *authController) Logout(ctx *fiber.Ctx) error {
	ssid := ctx.Cookies(constant.SESSION_COOKIE_NAME)

	cookie, err := auth.authUsecase.Logout(ssid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	ctx.Cookie(cookie)
	return ctx.Redirect("http://localhost:3000")
}
