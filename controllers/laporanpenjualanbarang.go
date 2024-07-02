package controllers

import (
	"Delimart/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LaporanPenjualanBarangController struct{}

func (controller *LaporanPenjualanBarangController) GetLaporanPenjualanBarang(c echo.Context) error {
	laporan, err := models.GetLaporanPenjualanBarang()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}
