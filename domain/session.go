package domain

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Session struct {
	ID         string    `json:"id" db:"id"`
	USER_ID    string    `json:"userId" db:"user_id"`
	IP_ADDRESS string    `json:"-" db:"ip_address"`
	USER_AGENT string    `json:"userAgent" db:"user_agent"`
	EXPIRED_AT time.Time `json:"expiredAt" db:"expired_at"`
	CREATED_AT time.Time `json:"createdAt" db:"created_at"`
}

type SessionRepository interface {
	Create(session *Session) error
	GetByID(id string) (*Session, error)
	DeleteById(id string) error
}

type SessionUsecase interface {
	Create(userId string, ipAddress string, userAgent string) (*fiber.Cookie, error)
	GetByID(id string) (*Session, error)
	DeleteById(id string) error
}
