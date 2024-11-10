package repository

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/jmoiron/sqlx"
)

type geoLocationRepository struct {
	db *sqlx.DB
}

func NewGeoLocationRepository(db *sqlx.DB) domain.GeoLocationRepository {
	return &geoLocationRepository{
		db: db,
	}
}

func (r *geoLocationRepository) GetProvinces() (*[]domain.Province, error) {
	provinces := make([]domain.Province, 0)
	err := r.db.Select(&provinces, "SELECT * FROM thai_provinces")
	if err != nil {
		return nil, err
	}
	return &provinces, nil
}

func (r *geoLocationRepository) GetDistrictsByProvinceId(provinceId string) (*[]domain.District, error) {
	districts := make([]domain.District, 0)
	err := r.db.Select(&districts, "SELECT * FROM thai_districts WHERE province_id = $1", provinceId)
	if err != nil {
		return nil, err
	}
	return &districts, nil
}

func (r *geoLocationRepository) GetSubDistrictsByDistrictId(districtId string) (*[]domain.SubDistrict, error) {
	subDistricts := make([]domain.SubDistrict, 0)
	err := r.db.Select(&subDistricts, "SELECT * FROM thai_subdistricts WHERE district_id = $1", districtId)
	if err != nil {
		return nil, err
	}
	return &subDistricts, nil
}
