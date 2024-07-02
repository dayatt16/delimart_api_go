package controllers

import (
	"Delimart/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LaporanArusStokController struct{}

func (controller *LaporanArusStokController) GetLaporanArusStok(c echo.Context) error {
	laproan, err := models.GetLaporanArusStok()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, laproan)
}
