package controllers

import (
	"Delimart/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LaporanDashboardController struct{}

func (controller *LaporanDashboardController) GetLaporanDashboard(c echo.Context) error {
	laporan, err := models.GetLaporanDashboard()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laporan)
}
