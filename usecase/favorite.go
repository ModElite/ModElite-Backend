package usecase

import "github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"

type favoriteUsecase struct {
	favoriteRepo   domain.FavoriteRepository
	productUsecase domain.ProductUsecase
}

func NewFavoriteUsecase(
	favoriteRepo domain.FavoriteRepository,
	productUsecase domain.ProductUsecase,
) domain.FavoriteUsecase {
	return &favoriteUsecase{
		favoriteRepo:   favoriteRepo,
		productUsecase: productUsecase,
	}
}

func (fu *favoriteUsecase) GetAll() (*[]domain.Favorite, error) {
	return fu.favoriteRepo.GetAll()
}

func (fu *favoriteUsecase) GetByID(id string) (*domain.Favorite, error) {
	favorite, err := fu.favoriteRepo.GetByID(id)
	if err != nil || favorite == nil {
		return nil, err
	}

	product, err := fu.productUsecase.GetProductWithOptionsAndSizes(favorite.PRODUCT_ID)
	if err != nil {
		return nil, err
	}
	return &domain.Favorite{
		PRODUCT:    *product,
		PRODUCT_ID: favorite.PRODUCT_ID,
	}, nil
}

func (fu *favoriteUsecase) GetByUserID(id string) (*[]domain.Favorite, error) {
	favorites, err := fu.favoriteRepo.GetByUserID(id)
	if err != nil {
		return nil, err
	}

	for i, favorite := range *favorites {
		product, err := fu.productUsecase.GetProductWithOptionsAndSizes(favorite.PRODUCT_ID)
		if err != nil {
			return nil, err
		}
		(*favorites)[i].PRODUCT = *product
	}
	return favorites, nil
}

func (fu *favoriteUsecase) Create(favorite *domain.Favorite) error {
	return fu.favoriteRepo.Create(favorite)
}

func (fu *favoriteUsecase) Delete(user_id string, product_id string) error {
	return fu.favoriteRepo.Delete(user_id, product_id)
}
