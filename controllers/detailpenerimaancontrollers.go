package controllers

import (
	"Delimart/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DetailPenerimaanController struct{}

func (controller *DetailPenerimaanController) CreateDetailPenerimaan(c echo.Context) error {
	var dp models.DetailPenerimaan
	if err := c.Bind(&dp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := models.CreateDetailPenerimaan(&dp); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, dp)
}
