package usecase

import (
	"time"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type sessionUsecase struct {
	sessionRepo domain.SessionRepository
}

func NewSessionUsecase(sessionRepo domain.SessionRepository) domain.SessionUsecase {
	return &sessionUsecase{
		sessionRepo: sessionRepo,
	}
}

func (su *sessionUsecase) Create(userId string, ipAddress string, userAgent string) (*fiber.Cookie, error) {
	session := &domain.Session{
		ID:         uuid.New().String(),
		USER_ID:    userId,
		IP_ADDRESS: ipAddress,
		USER_AGENT: userAgent,
		EXPIRED_AT: time.Now().Add(time.Hour * 24 * 3),
		CREATED_AT: time.Now(),
	}
	if err := su.sessionRepo.Create(session); err != nil {
		return nil, err
	}
	cookie := &fiber.Cookie{
		Name:     constant.SESSION_COOKIE_NAME,
		Value:    session.ID,
		Expires:  time.Now().Add(time.Hour * 24 * 3),
		HTTPOnly: true,
		Secure:   false,
		Domain:   "sssboom.xyz",
		Path:     "/",
		SameSite: "None",
	}

	return cookie, nil
}

func (su *sessionUsecase) GetByID(id string) (*domain.Session, error) {
	return su.sessionRepo.GetByID(id)
}

func (su *sessionUsecase) DeleteById(id string) error {
	return su.sessionRepo.DeleteById(id)
}
