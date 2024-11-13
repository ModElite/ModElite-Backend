package controller

import (
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type geoLocationController struct {
	validator          domain.ValidatorUsecase
	geoLocationUsecase domain.GeoLocationUsecase
}

func NewGeoLocationController(
	validator domain.ValidatorUsecase,
	geoLocationUsecase domain.GeoLocationUsecase,
) *geoLocationController {
	return &geoLocationController{
		validator:          validator,
		geoLocationUsecase: geoLocationUsecase,
	}
}

// GetProvinces godoc
// @Summary Get all provinces
// @Description Get all provinces
// @Tags GeoLocation
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response{data=[]domain.Province}
// @Router /api/geo-location/provinces [get]
func (c *geoLocationController) GetProvinces(ctx *fiber.Ctx) error {
	provinces, err := c.geoLocationUsecase.GetProvinces()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    provinces,
	})
}

// GetDistrictsByProvinceId godoc
// @Summary Get all districts by province id
// @Description Get all districts by province id
// @Tags GeoLocation
// @Accept json
// @Produce json
// @Param province_id path string true "Province ID"
// @Success 200 {object} domain.Response{data=[]domain.District}
// @Router /api/geo-location/districts/{province_id} [get]
func (c *geoLocationController) GetDistrictsByProvinceId(ctx *fiber.Ctx) error {
	provinceId := ctx.Params("province_id")
	if provinceId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	} else if !c.validator.ValidateUUID(provinceId) {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}
	districts, err := c.geoLocationUsecase.GetDistrictsByProvinceId(provinceId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    districts,
	})
}

// GetSubDistrictsByDistrictId godoc
// @Summary Get all sub-districts by district id
// @Description Get all sub-districts by district id
// @Tags GeoLocation
// @Accept json
// @Produce json
// @Param district_id path string true "District ID"
// @Success 200 {object} domain.Response{data=[]domain.SubDistrict}
// @Router /api/geo-location/sub-districts/{district_id} [get]
func (c *geoLocationController) GetSubDistrictsByDistrictId(ctx *fiber.Ctx) error {
	districtId := ctx.Params("district_id")
	if districtId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	} else if !c.validator.ValidateUUID(districtId) {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
			SUCCESS: false,
		})
	}
	subDistricts, err := c.geoLocationUsecase.GetSubDistrictsByDistrictId(districtId)
	if err != nil {
		println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
			SUCCESS: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		MESSAGE: constant.MESSAGE_SUCCESS,
		SUCCESS: true,
		DATA:    subDistricts,
	})
}
