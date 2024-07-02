package controllers

import (
	"Delimart/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BarangController struct{}

func (controller *BarangController) GetAllBarang(c echo.Context) error {
	barangs, err := models.GetAllBarang()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, barangs)
}

func (controller *BarangController) GetBarangByKodeBarang(c echo.Context) error {
	kdBarang := c.Param("kd_barang")
	barang, err := models.GetBarangByKodeBarang(kdBarang)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, barang)
}

func (controller *BarangController) GetBarangByKodeSupp(c echo.Context) error {
	kdSupplier := c.Param("kd_supplier")
	barangs, err := models.GetBarangByKodeSupp(kdSupplier)
	if err != nil {
		if err.Error() == "barang not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Barang tidak ditemukan"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal memuat barang"})
	}
	return c.JSON(http.StatusOK, barangs)
}

func (controller *BarangController) SearchBarang(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Query parameter 'quary' is required"})
	}

	hasil, err := models.SearchBarang(query)
	if err != nil {
		log.Printf("Error searching pegawai: %v\n", err)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "No matching barang found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search barang"})
	}

	return c.JSON(http.StatusOK, hasil)
}

func (controller *BarangController) GetBarangByNama(c echo.Context) error {
	nama := c.Param("nama")
	barang, err := models.GetBarangByNama(nama)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, barang)
}

func (controller *BarangController) CreateBarang(c echo.Context) error {
	var barang models.Barang
	if err := c.Bind(&barang); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := models.CreateBarang(&barang); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, barang)
}

func (controller *BarangController) UpdateBarang(c echo.Context) error {
	kdBarang := c.Param("kd_barang")
	var barang models.Barang
	if err := c.Bind(&barang); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updatedBarang, err := models.UpdateBarangByKodeBarang(kdBarang, &barang)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedBarang)
}

func (controller *BarangController) DeleteBarang(c echo.Context) error {
	kdBarang := c.Param("kd_barang")
	if err := models.DeleteBarangByKodeBarang(kdBarang); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Barang successfully deleted"})
}
