package usecase

import "github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"

type geoLocationUseCase struct {
	geoLocationRepository domain.GeoLocationRepository
}

func NewGeoLocationUseCase(geoLocationRepository domain.GeoLocationRepository) domain.GeoLocationUsecase {
	return &geoLocationUseCase{
		geoLocationRepository: geoLocationRepository,
	}
}

func (g *geoLocationUseCase) GetProvinces() (*[]domain.Province, error) {
	return g.geoLocationRepository.GetProvinces()
}

func (g *geoLocationUseCase) GetDistrictsByProvinceId(provinceId string) (*[]domain.District, error) {
	return g.geoLocationRepository.GetDistrictsByProvinceId(provinceId)
}

func (g *geoLocationUseCase) GetSubDistrictsByDistrictId(districtId string) (*[]domain.SubDistrict, error) {
	return g.geoLocationRepository.GetSubDistrictsByDistrictId(districtId)
}
