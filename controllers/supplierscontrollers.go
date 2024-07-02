package controllers

import (
	"Delimart/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SupplierController struct{}

func (controller *SupplierController) GetAllSuppliers(c echo.Context) error {
	suppliers, err := models.GetAllSuppliers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, suppliers)
}

func (controller *SupplierController) GetSupplierByKodeSupplier(c echo.Context) error {
	kdSupplier := c.Param("kd_supplier")
	supplier, err := models.GetSupplierByKodeSupplier(kdSupplier)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, supplier)
}

func (controller *SupplierController) GetKodeSupplierByNama(c echo.Context) error {
	nama := c.Param("nama")
	supplier, err := models.GetKodeSupplierByNama(nama)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, supplier)
}

func (controller *SupplierController) SearchSupplier(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Query parameter 'query' is required"})
	}

	hasil, err := models.SearchSupplier(query)
	if err != nil {
		log.Printf("Error searching supplier: %v\n", err)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "No matching supplier found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search supplier"})
	}

	return c.JSON(http.StatusOK, hasil)
}

func (controller *SupplierController) CreateSupplier(c echo.Context) error {
	var supplier models.Supplier
	if err := c.Bind(&supplier); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := models.CreateSupplier(&supplier); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, supplier)
}

func (controller *SupplierController) UpdateSupplier(c echo.Context) error {
	kdSupplier := c.Param("kd_supplier")
	var supplier models.Supplier
	if err := c.Bind(&supplier); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updatedSupplier, err := models.UpdateSupplierByKodeSupplier(kdSupplier, &supplier)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedSupplier)
}

func (controller *SupplierController) DeleteSupplier(c echo.Context) error {
	kdSupplier := c.Param("kd_supplier")
	if err := models.DeleteSupplierByKodeSupplier(kdSupplier); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Supplier successfully deleted"})
}
