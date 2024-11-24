package usecase

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/google/uuid"
)

type sizeUsecase struct {
	sizeRepo domain.SizeRepository
}

func NewSizeUsecase(sizeRepo domain.SizeRepository) domain.SizeUsecase {
	return &sizeUsecase{
		sizeRepo: sizeRepo,
	}
}

func (su *sizeUsecase) GetAll() (*[]domain.Size, error) {
	return su.sizeRepo.GetAll()
}

func (su *sizeUsecase) GetByID(id string) (*domain.Size, error) {
	return su.sizeRepo.GetByID(id)
}

func (su *sizeUsecase) Create(size *domain.Size) error {
	return su.sizeRepo.Create(&domain.Size{
		ID:   uuid.New().String(),
		SIZE: size.SIZE,
	})
}

func (su *sizeUsecase) Update(size *domain.Size) error {
	return su.sizeRepo.Update(size)
}

func (su *sizeUsecase) Delete(id string) error {
	return su.sizeRepo.Delete(id)
}
