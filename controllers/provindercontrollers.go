package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "Delimart/models"
)

type ProviderController struct{}

func (controller *ProviderController) GetAllProviders(c echo.Context) error {
    providers, err := models.GetAllProviders()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, providers)
}

func (controller *ProviderController) GetProviderByKodeProvider(c echo.Context) error {
    kdProvider := c.Param("kd_provider")
    provider, err := models.GetProviderByKodeProvider(kdProvider)
    if err != nil {
        if err.Error() == "Provider not found" {
            return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, provider)
}

func (controller *ProviderController) CreateProvider(c echo.Context) error {
    var provider models.Provider
    if err := c.Bind(&provider); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    if err := models.CreateProvider(&provider); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, provider)
}

func (controller *ProviderController) UpdateProvider(c echo.Context) error {
    kdProvider := c.Param("kd_provider")
    var provider models.Provider
    if err := c.Bind(&provider); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    updatedProvider, err := models.UpdateProviderByKodeProvider(kdProvider, &provider)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, updatedProvider)
}

func (controller *ProviderController) DeleteProvider(c echo.Context) error {
    kdProvider := c.Param("kd_provider")
    if err := models.DeleteProviderByKodeProvider(kdProvider); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Provider successfully deleted"})
}
