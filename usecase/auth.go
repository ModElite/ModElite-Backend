package usecase

import (
	"fmt"
	"time"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type authUsecase struct {
	googleUsecase  domain.GoogleUsecase
	userUsecase    domain.UserUsecase
	sessionUsecase domain.SessionUsecase
}

func NewAuthUsecase(
	googleUsecase domain.GoogleUsecase,
	userUsecase domain.UserUsecase,
	sessionUsecase domain.SessionUsecase,
) domain.AuthUsecase {
	return &authUsecase{
		googleUsecase:  googleUsecase,
		userUsecase:    userUsecase,
		sessionUsecase: sessionUsecase,
	}
}

func (u *authUsecase) SignInWithGoogle(c *fiber.Ctx) (*fiber.Cookie, error) {
	token, err := u.googleUsecase.GetToken(c)
	if err != nil {
		return nil, err
	}

	profile, err := u.googleUsecase.GetProfile(token.AccessToken)
	if err != nil {
		return nil, err
	}

	user, err := u.userUsecase.GetByEmail(profile.EMAIL)
	if err != nil {
		return nil, err
	}
	if user == nil {
		user, err = u.userUsecase.CreateFromGoogle(profile.NAME, profile.EMAIL, profile.ID, profile.PICTURE)
		if err != nil {
			return nil, err
		}
	}

	cookie, err := u.sessionUsecase.Create(user.ID, c.IP(), c.Get("User-Agent"))
	if err != nil {
		return nil, fmt.Errorf("cannot create session to sign in with google %w", err)
	}
	return cookie, nil
}

func (u *authUsecase) Logout(ssid string) (*fiber.Cookie, error) {
	session, err := u.sessionUsecase.GetByID(ssid)
	if err != nil {
		return nil, fmt.Errorf("cannot get session to logout %w", err)
	}

	err = u.sessionUsecase.DeleteById(session.ID)
	if err != nil {
		return nil, fmt.Errorf("cannot delete session to logout %w", err)
	}
	cookie := &fiber.Cookie{
		Name:     constant.SESSION_COOKIE_NAME,
		Value:    "",
		Expires:  time.Unix(0, 0),
		HTTPOnly: true,
	}
	return cookie, nil
}
