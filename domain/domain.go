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
	FavoriteRepository      FavoriteRepository
	TagsRepository          TagsRepository
	ProductTagsRepository   ProductTagsRepository
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
	FavoriteUsecase      FavoriteUsecase
	TagsUsecase          TagsUsecase
}
