package controllers

import (
	"Delimart/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PegawaiController struct{}

func (controller *PegawaiController) GetAllPegawai(c echo.Context) error {
	pegawai, err := models.GetAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, pegawai)
}

func (controller *PegawaiController) GetPegawaiByKodePegawai(c echo.Context) error {
	kdPegawai := c.Param("kd_pegawai")
	pegawai, err := models.GetPegawaiByKodePegawai(kdPegawai)
	if err != nil {
		if err.Error() == "pegawai not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, pegawai)
}

func (controller *PegawaiController) SearchPegawai(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Query parameter 'quary' is required"})
	}

	hasil, err := models.SearchPegawai(query)
	if err != nil {
		log.Printf("Error searching pegawai: %v\n", err)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "No matching pegawai found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search pegawai"})
	}

	return c.JSON(http.StatusOK, hasil)
}

func (controller *PegawaiController) CreatePegawai(c echo.Context) error {
	var pegawai models.Pegawai
	if err := c.Bind(&pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := models.CreatePegawai(&pegawai); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, pegawai)
}

func (controller *PegawaiController) UpdatePegawai(c echo.Context) error {
	kdPegawai := c.Param("kd_pegawai")
	var pegawai models.Pegawai
	if err := c.Bind(&pegawai); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := models.UpdatePegawaiByKodePegawai(kdPegawai, &pegawai); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, pegawai)
}

func (controller *PegawaiController) DeletePegawai(c echo.Context) error {
	kdPegawai := c.Param("kd_pegawai")
	if err := models.DeletePegawaiByKodePegawai(kdPegawai); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Pegawai successfully deleted"})
}
