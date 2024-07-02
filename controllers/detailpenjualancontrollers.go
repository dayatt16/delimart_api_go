package controllers

import (
	"Delimart/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DetailPenjualanController struct{}

func (controller *DetailPenjualanController) CreateDetailPenjualan(c echo.Context) error {
	var dpenj models.DetailPenjualan
	if err := c.Bind(&dpenj); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := models.CreateDetailPenjualan(&dpenj); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, dpenj)
}
