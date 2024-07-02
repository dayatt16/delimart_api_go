package controllers

import (
	"Delimart/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LaporanPenjualanHomeController struct{}

func (controller *LaporanPenjualanHomeController) GetLaporanPenjualanHome(c echo.Context) error {
	laporan, err := models.GetLaporanPenjualanHome()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}
