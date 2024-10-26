package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/google/uuid"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) CheckAdmin(id string) (bool, error) {
	user, err := u.userRepository.Get(id)
	if err != nil || user == nil {
		return false, fmt.Errorf("cannot get user by id %s", id)
	}
	if user.ROLE != domain.AdminAccount {
		return false, nil
	}
	return true, nil
}

func (u *userUsecase) CreateFromGoogle(name string, email string, google_id string) (*domain.User, error) {
	FirstName := ""
	LastName := ""
	result := strings.Split(name, " ")
	if len(result) == 2 {
		FirstName = result[0]
		LastName = result[1]
	}

	user := &domain.User{
		ID:          uuid.New().String(),
		EMAIL:       email,
		GOOGLE_ID:   google_id,
		FIRST_NAME:  FirstName,
		LAST_NAME:   LastName,
		PHONE:       "",
		PROFILE_URL: "",
		UPDATED_AT:  time.Now(),
		CREATED_AT:  time.Now(),
	}

	if err := u.userRepository.Create(user); err != nil {
		return nil, fmt.Errorf("cannot create user from google auth email %s", email)
	}
	return user, nil
}

func (u *userUsecase) Get(id string) (*domain.User, error) {
	user, err := u.userRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("cannot get user by id %s", id)
	}
	return user, nil
}

func (u *userUsecase) GetByEmail(email string) (*domain.User, error) {
	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("cannot get user by email %s", email)
	}
	return user, nil
}
