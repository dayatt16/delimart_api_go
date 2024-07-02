package controllers

import (
	"net/http"

	"Delimart/models"

	"github.com/labstack/echo/v4"
)

type PrefixController struct{}

func (controller *PrefixController) GetAllPrefix(c echo.Context) error {
	prefixes, err := models.GetAllPrefix()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, prefixes)
}

func (controller *PrefixController) GetPrefixByKodePrefix(c echo.Context) error {
	kdPrefix := c.Param("kd_prefix")
	prefix, err := models.GetPrefixByKodePrefix(kdPrefix)
	if err != nil {
		if err.Error() == "Prefix not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, prefix)
}

func (controller *PrefixController) GetProviderByPrefix(c echo.Context) error {
	prefix := c.Param("prefix")
	barang, err := models.GetProviderByPrefix(prefix)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, barang)
}

func (controller *PrefixController) CreatePrefix(c echo.Context) error {
	var prefix models.Prefix
	if err := c.Bind(&prefix); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := models.CreatePrefix(&prefix); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, prefix)
}

func (controller *PrefixController) UpdatePrefix(c echo.Context) error {
	kdPrefix := c.Param("kd_prefix")
	var prefix models.Prefix
	if err := c.Bind(&prefix); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updatedPrefix, err := models.UpdatePrefixByKodePrefix(kdPrefix, &prefix)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedPrefix)
}

func (controller *PrefixController) DeletePrefix(c echo.Context) error {
	kdPrefix := c.Param("kd_prefix")
	if err := models.DeletePrefixByKodePrefix(kdPrefix); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Prefix successfully deleted"})
}
