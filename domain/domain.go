package domain

type Repository struct {
	UserRepository    UserRepository
	SessionRepository SessionRepository
	SellerRepository  SellerRepository
}

type Usecase struct {
	AuthUsecase    AuthUsecase
	GoogleUsecase  GoogleUsecase
	UserUsecase    UserUsecase
	SessionUsecase SessionUsecase
	SellerUsecase  SellerUsecase
}
