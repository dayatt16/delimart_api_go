package controllers

import (
	"Delimart/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PulsaController struct{}

func (controller *PulsaController) GetAllPulsa(c echo.Context) error {
	pulsa, err := models.GetAllPulsa()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, pulsa)
}

func (controller *PulsaController) GetPulsaByKodePulsa(c echo.Context) error {
	kdPulsa := c.Param("kd_pulsa")
	pulsa, err := models.GetPulsaByKodePulsa(kdPulsa)
	if err != nil {
		if err.Error() == "Pulsa not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, pulsa)
}

func (controller *PulsaController) SearchPulsa(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Query parameter 'query' is required"})
	}

	hasil, err := models.SearchPulsa(query)
	if err != nil {
		log.Printf("Error searching pulsa: %v\n", err)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "No matching pulsa found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search pulsa"})
	}

	return c.JSON(http.StatusOK, hasil)
}

func (controller *PulsaController) GetPulsaByProvider(c echo.Context) error {
	kdPulsa := c.Param("provider")
	pulsa, err := models.GetPulsaByProvider(kdPulsa)
	if err != nil {
		if err.Error() == "Pulsa not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, pulsa)
}

func (controller *PulsaController) CreatePulsa(c echo.Context) error {
	var pulsa models.Pulsa
	if err := c.Bind(&pulsa); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := models.CreatePulsa(&pulsa); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, pulsa)
}

func (controller *PulsaController) UpdatePulsa(c echo.Context) error {
	kdPulsa := c.Param("kd_pulsa")
	var pulsa models.Pulsa
	if err := c.Bind(&pulsa); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updatedPulsa, err := models.UpdatePulsaByKodePulsa(kdPulsa, &pulsa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedPulsa)
}

func (controller *PulsaController) DeletePulsa(c echo.Context) error {
	kdPulsa := c.Param("kd_pulsa")
	if err := models.DeletePulsaByKodePulsa(kdPulsa); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Pulsa successfully deleted"})
}
