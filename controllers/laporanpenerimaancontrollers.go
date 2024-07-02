package controllers

import (
	"Delimart/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LaporanPenerimaanHarianController struct{}
type LaporanPenerimaanBulananController struct{}
type LaporanPenerimaanTahunanController struct{}
type LaporanPenerimaanAllController struct{}

func (controller *LaporanPenerimaanHarianController) GetLaporanPenerimaanHarian(c echo.Context) error {
	laporan, err := models.GetLaporanPenerimaanHarian()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}

func (controller *LaporanPenerimaanBulananController) GetLaporanPenerimaanBulanan(c echo.Context) error {
	laporan, err := models.GetLaporanPenerimaanBulanan()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}

func (controller *LaporanPenerimaanTahunanController) GetLaporanPenerimaanTahunan(c echo.Context) error {
	laporan, err := models.GetLaporanPenerimaanTahunan()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}

func (controller *LaporanPenerimaanAllController) GetLaporanPenerimaanAll(c echo.Context) error {
	laporan, err := models.GetLaporanPenerimaanAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}
