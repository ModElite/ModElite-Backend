package domain

import (
	"time"
)

type AccountType string

const (
	AdminAccount AccountType = "ADMIN"
	UserAccount  AccountType = "USER"
)

type User struct {
	ID          string      `json:"id,omitempty" db:"id"`
	EMAIL       string      `json:"email" db:"email"`
	GOOGLE_ID   string      `json:"google_id" db:"google_id"`
	FIRST_NAME  string      `json:"firstName" db:"first_name"`
	LAST_NAME   string      `json:"lastName" db:"last_name"`
	PHONE       string      `json:"phone" db:"phone"`
	PROFILE_URL string      `json:"profileUrl" db:"profile_url"`
	ROLE        AccountType `json:"role" db:"role"`
	UPDATED_AT  time.Time   `json:"updateAt" db:"updated_at"`
	CREATED_AT  time.Time   `json:"createdAt" db:"created_at"`
}

type UserRepository interface {
	Create(*User) (err error)
	Get(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(*User) (err error)
}

type UserUsecase interface {
	CheckAdmin(id string) (bool, error)
	CreateFromGoogle(name string, email string, google_id string) (*User, error)
	Get(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(userId string, userUpdate *User) (err error)
}
