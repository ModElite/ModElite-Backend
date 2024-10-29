package domain

type Repository struct {
	UserRepository    UserRepository
	SessionRepository SessionRepository
	SellerRepository  SellerRepository
	ProductRepository ProductRepository
}

type Usecase struct {
	AuthUsecase    AuthUsecase
	GoogleUsecase  GoogleUsecase
	UserUsecase    UserUsecase
	SessionUsecase SessionUsecase
	SellerUsecase  SellerUsecase
	ProductUsecase ProductUsecase
}
