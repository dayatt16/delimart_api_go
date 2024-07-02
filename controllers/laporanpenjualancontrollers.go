package controllers

import (
	"Delimart/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LaporanPenjualanHarianController struct{}
type LaporanPenjualanBulananController struct{}
type LaporanPenjualanTahunanController struct{}
type LaporanPenjualanAllController struct{}

func (controller *LaporanPenjualanHarianController) GetLaporanPenjualanHarian(c echo.Context) error {
	laporan, err := models.GetLaporanPenjualanHarian()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}

func (controller *LaporanPenjualanBulananController) GetLaporanPenjualanBulanan(c echo.Context) error {
	laporan, err := models.GetLaporanPenjualanBulanan()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}
func (controller *LaporanPenjualanTahunanController) GetLaporanPenjualanTahunan(c echo.Context) error {
	laporan, err := models.GetLaporanPenjualanTahunan()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}

func (controller *LaporanPenjualanAllController) GetLaporanPenjualanAll(c echo.Context) error {
	laporan, err := models.GetLaporanPenjualanAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}
