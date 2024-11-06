package domain

type Repository struct {
	UserRepository          UserRepository
	SessionRepository       SessionRepository
	SellerRepository        SellerRepository
	ProductRepository       ProductRepository
	ProductOptionRepository ProductOptionRepository
	ProductSizeRepository   ProductSizeRepository
	SizeRepository          SizeRepository
	AddressRepository       AddressRepository
}

type Usecase struct {
	AuthUsecase          AuthUsecase
	GoogleUsecase        GoogleUsecase
	UserUsecase          UserUsecase
	SessionUsecase       SessionUsecase
	SellerUsecase        SellerUsecase
	ProductUsecase       ProductUsecase
	ProductOptionUsecase ProductOptionUsecase
	ProductSizeUsecase   ProductSizeUsecase
	SizeUsecase          SizeUsecase
	AddressUsecase       AddressUsecase
}
