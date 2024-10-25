package domain

type Repository struct {
	UserRepository    UserRepository
	SessionRepository SessionRepository
}

type Usecase struct {
	AuthUsecase    AuthUsecase
	GoogleUsecase  GoogleUsecase
	UserUsecase    UserUsecase
	SessionUsecase SessionUsecase
}
