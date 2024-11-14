package usecase

import (
	"fmt"
	"sync"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/google/uuid"
)

type productUsecase struct {
	productRepo          domain.ProductRepository
	productOptionUsecase domain.ProductOptionUsecase
	productSizeUsecase   domain.ProductSizeUsecase
	tagUsecase           domain.TagUsecase
	userUsecase          domain.UserUsecase
	sellerUsecase        domain.SellerUsecase
}

func NewProductUsecase(
	productRepo domain.ProductRepository,
	productOptionUsecase domain.ProductOptionUsecase,
	productSizeUsecase domain.ProductSizeUsecase,
	tagUsecase domain.TagUsecase,
	userUsecase domain.UserUsecase,
	sellerUsecase domain.SellerUsecase,
) domain.ProductUsecase {
	return &productUsecase{
		productRepo:          productRepo,
		productOptionUsecase: productOptionUsecase,
		productSizeUsecase:   productSizeUsecase,
		tagUsecase:           tagUsecase,
		userUsecase:          userUsecase,
		sellerUsecase:        sellerUsecase,
	}
}

func (u *productUsecase) CheckPermissionCanModifyProduct(ownerID string, productID string) (bool, error) {
	product, err := u.productRepo.GetByID(productID)
	if err != nil || product == nil {
		return false, fmt.Errorf("error product getbyid: %w", err)
	}

	seller, err := u.sellerUsecase.GetByID(product.SELLER_ID)
	if err != nil || seller == nil {
		return false, fmt.Errorf("error seller getbyid: %w", err)
	} else if seller.OWNER_ID != ownerID {
		if isAdmin, err := u.userUsecase.CheckAdmin(ownerID); err != nil {
			return false, fmt.Errorf("error user checkadmin: %w", err)
		} else if !isAdmin {
			return false, nil
		}
	}

	return true, nil
}

func (u *productUsecase) GetAll() (*[]domain.Product, error) {
	products, err := u.productRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	return products, nil
}

func (u *productUsecase) GetAllProductWithOptionsAndSizes() (*[]domain.Product, error) {
	products, err := u.productRepo.GetAllProductWithOptionsAndSizes()
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	for i, product := range *products {
		productTags, err := u.tagUsecase.GetTagByProductID(product.ID)
		if err != nil {
			return nil, fmt.Errorf("error product getall: %w", err)
		}
		(*products)[i].TAGS = productTags
	}

	return products, nil
}

func (u *productUsecase) GetProductWithOptionsAndSizes(productId string) (*domain.Product, error) {
	product, err := u.productRepo.GetProductWithOptionsAndSizes(productId)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	} else if product == nil {
		return nil, nil
	}

	productTags, err := u.tagUsecase.GetTagByProductID(product.ID)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}
	product.TAGS = productTags

	return product, nil
}

func (u *productUsecase) GetProductsBySeller(sellerID string) (*[]domain.Product, error) {
	products, err := u.productRepo.GetProductsBySeller(sellerID)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	for i, product := range *products {
		productTags, err := u.tagUsecase.GetTagByProductID(product.ID)
		if err != nil {
			return nil, fmt.Errorf("error product getall: %w", err)
		}
		(*products)[i].TAGS = productTags
	}

	return products, nil
}

func (u *productUsecase) GetByID(id string) (*domain.Product, error) {
	product, err := u.productRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	return product, nil
}

func (u *productUsecase) GetBySellerID(SellerID string) (*[]domain.Product, error) {
	products, err := u.productRepo.GetBySellerID(SellerID)
	if err != nil {
		return nil, fmt.Errorf("error product getall: %w", err)
	}

	return products, nil
}

func (u *productUsecase) Create(product *domain.Product) (*string, error) {
	product.ID = uuid.New().String()

	if err := u.productRepo.Create(&domain.Product{
		ID:          product.ID,
		SELLER_ID:   product.SELLER_ID,
		NAME:        product.NAME,
		DESCRIPTION: product.DESCRIPTION,
		PRICE:       product.PRICE,
		IMAGE_URL:   product.IMAGE_URL,
		STATUS:      string(domain.ProductActive),
	}); err != nil {
		return nil, fmt.Errorf("error product create: %w", err)
	}

	for _, option := range *product.PRODUCT_OPTION {
		option.ID = uuid.New().String()
		option.PRODUCT_ID = product.ID
		if err := u.productOptionUsecase.Create(&domain.ProductOption{
			ID:         option.ID,
			PRODUCT_ID: option.PRODUCT_ID,
			LABEL:      option.LABEL,
			IMAGE_URL:  option.IMAGE_URL,
		}); err != nil {
			return nil, fmt.Errorf("error product option create: %w", err)
		}

		for _, size := range *option.PRODUCT_SIZE {
			size.ID = uuid.New().String()
			size.PRODUCT_OPTION_ID = option.ID
			if err := u.productSizeUsecase.Create(&domain.ProductSize{
				ID:                size.ID,
				PRODUCT_OPTION_ID: size.PRODUCT_OPTION_ID,
				SIZE_ID:           size.SIZE_ID,
				QUANTITY:          size.QUANTITY,
			}); err != nil {
				return nil, fmt.Errorf("error product size create: %w", err)
			}
		}
	}

	return &product.ID, nil
}

func (u *productUsecase) Update(newProduct *domain.Product) error {
	oldProductOption, err := u.productOptionUsecase.GetByProductIDAndFilterActive(newProduct.ID)
	if err != nil {
		return fmt.Errorf("error product option getbyproductid: %w", err)
	}

	// Delete old product option and size if not exist in new product
	var wg sync.WaitGroup
	wg.Add(len(*oldProductOption))
	for _, option := range *oldProductOption {
		go func(option domain.ProductOption) {
			defer wg.Done()
			found := false
			for _, newOption := range *newProduct.PRODUCT_OPTION {
				if option.ID == newOption.ID && newOption.ID != "" {
					found = true
					break
				}
			}

			if !found {
				if err := u.productOptionUsecase.SoftDeleteProductOptionAndSizeByProductID(option.ID); err != nil {
					fmt.Printf("error product option softdelete: %v\n", err)
				}

			}
		}(option)
	}
	wg.Wait()

	// Update product
	if err := u.productRepo.Update(&domain.Product{
		ID:          newProduct.ID,
		NAME:        newProduct.NAME,
		DESCRIPTION: newProduct.DESCRIPTION,
		PRICE:       newProduct.PRICE,
		IMAGE_URL:   newProduct.IMAGE_URL,
		STATUS:      string(domain.ProductActive),
	}); err != nil {
		return fmt.Errorf("error product update: %w", err)
	}

	// Create or update product option and size in new product
	for _, option := range *newProduct.PRODUCT_OPTION {
		if option.ID == "" {
			option.ID = uuid.New().String()
			option.PRODUCT_ID = newProduct.ID
			if err := u.productOptionUsecase.Create(&domain.ProductOption{
				ID:         option.ID,
				PRODUCT_ID: option.PRODUCT_ID,
				LABEL:      option.LABEL,
				IMAGE_URL:  option.IMAGE_URL,
			}); err != nil {
				return fmt.Errorf("error product option create: %w", err)
			}

			for _, size := range *option.PRODUCT_SIZE {
				size.ID = uuid.New().String()
				size.PRODUCT_OPTION_ID = option.ID
				if err := u.productSizeUsecase.Create(&domain.ProductSize{
					ID:                size.ID,
					PRODUCT_OPTION_ID: size.PRODUCT_OPTION_ID,
					SIZE_ID:           size.SIZE_ID,
					QUANTITY:          size.QUANTITY,
				}); err != nil {
					return fmt.Errorf("error product size create: %w", err)
				}
			}
		} else {
			if err := u.productOptionUsecase.Update(&domain.ProductOption{
				ID:        option.ID,
				LABEL:     option.LABEL,
				IMAGE_URL: option.IMAGE_URL,
			}); err != nil {
				return fmt.Errorf("error product option update: %w", err)
			}

			for _, size := range *option.PRODUCT_SIZE {
				if size.ID == "" {
					size.ID = uuid.New().String()
					size.PRODUCT_OPTION_ID = option.ID
					if err := u.productSizeUsecase.Create(&domain.ProductSize{
						ID:                size.ID,
						PRODUCT_OPTION_ID: size.PRODUCT_OPTION_ID,
						SIZE_ID:           size.SIZE_ID,
						QUANTITY:          size.QUANTITY,
					}); err != nil {
						return fmt.Errorf("error product size create: %w", err)
					}
				} else {
					if err := u.productSizeUsecase.Update(&domain.ProductSize{
						ID:       size.ID,
						QUANTITY: size.QUANTITY,
					}); err != nil {
						return fmt.Errorf("error product size update: %w", err)
					}
				}
			}
		}
	}

	return nil
}

func (u *productUsecase) SoftDeleteWithOptionsAndSizes(id string) error {
	if err := u.productRepo.SoftDeleteWithOptionsAndSizes(id); err != nil {
		return fmt.Errorf("error product softdelete: %w", err)
	}

	return nil
}
