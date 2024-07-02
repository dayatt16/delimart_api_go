package controllers

import (
	"Delimart/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KategoriController struct{}

func (controller *KategoriController) GetAllKategori(c echo.Context) error {
    kategoris, err := models.GetAllKategori()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, kategoris)
}

func (controller *KategoriController) GetKategoriByKodeKategori(c echo.Context) error {
    kdKategori := c.Param("kd_kategori")
    kategori, err := models.GetKategoriByKodeKategori(kdKategori)
    if err != nil {
        if err.Error() == "Kategori not found" {
            return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, kategori)
}

func (controller *KategoriController) SearchKategori(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Query parameter 'query' is required"})
	}

	hasil, err := models.SearchKategori(query)
	if err != nil {
		log.Printf("Error searching kategori: %v\n", err)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "No matching kategori found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search kategori"})
	}

	return c.JSON(http.StatusOK, hasil)
}

func (controller *KategoriController) CreateKategori(c echo.Context) error {
    var kategori models.Kategori
    if err := c.Bind(&kategori); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    if err := models.CreateKategori(&kategori); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, kategori)
}

func (controller *KategoriController) UpdateKategori(c echo.Context) error {
    kdKategori := c.Param("kd_kategori")
    var kategori models.Kategori
    if err := c.Bind(&kategori); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    updatedKategori, err := models.UpdateKategoriByKodeKategori(kdKategori, &kategori)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, updatedKategori)
}

func (controller *KategoriController) DeleteKategori(c echo.Context) error {
    kdKategori := c.Param("kd_kategori")
    if err := models.DeleteKategoriByKodeKategori(kdKategori); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Kategori successfully deleted"})
}
