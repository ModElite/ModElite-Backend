package domain

type Province struct {
	ID      string `json:"id" db:"id"`
	NAME_TH string `json:"nameTh" db:"name_th"`
}

type District struct {
	ID          string `json:"id" db:"id"`
	NAME_TH     string `json:"nameTh" db:"name_th"`
	PROVINCE_ID string `json:"provinceId" db:"province_id"`
}

type SubDistrict struct {
	ID          string `json:"id" db:"id"`
	NAME_TH     string `json:"nameTh" db:"name_th"`
	DISTRICT_ID string `json:"districtId" db:"district_id"`
	ZIPCODE     string `json:"zipcode" db:"zip_code"`
}

type GeoLocationRepository interface {
	GetProvinces() (*[]Province, error)
	GetDistrictsByProvinceId(provinceId string) (*[]District, error)
	GetSubDistrictsByDistrictId(districtId string) (*[]SubDistrict, error)
}

type GeoLocationUsecase interface {
	GetProvinces() (*[]Province, error)
	GetDistrictsByProvinceId(provinceId string) (*[]District, error)
	GetSubDistrictsByDistrictId(districtId string) (*[]SubDistrict, error)
}
