package domain

type Repository struct {
}

type Usecase struct {
	AuthUsecase   AuthUsecase
	GoogleUsecase GoogleUsecase
}
