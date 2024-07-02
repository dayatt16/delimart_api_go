package controllers

import (
	"Delimart/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PenerimaanController struct{}

func (controller *PenerimaanController) GetNotaAuto(c echo.Context) error {
	nota, err := models.GetNotaAuto()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// Mengembalikan respons JSON dengan urutan sebagai string
	return c.JSON(http.StatusOK, map[string]string{"nota": nota})
}

//	func (controller *PenerimaanController) GetPenerimaanByKodeTransaksiTerima(c echo.Context) error {
//		kdTransaksiTerima := c.Param("kd_transaksi_terima")
//		penerimaan, err := models.GetPenerimaanByKodeTransaksiTerima(kdTransaksiTerima)
//		if err != nil {
//			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
//		}
//		return c.JSON(http.StatusOK, penerimaan)
//	}
func (controller *PenerimaanController) GetPenerimaanByNota(c echo.Context) error {
	nota := c.Param("nota")
	penerimaans, err := models.GetPenerimaanByNota(nota)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, penerimaans)
}

func (controller *PenerimaanController) CreatePenerimaan(c echo.Context) error {
	var penerimaan models.Penerimaan
	if err := c.Bind(&penerimaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Input tidak valid"})
	}

	if penerimaan.JumlahBarangDiterima < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Minimal barang masuk 1"})
	}

	err := models.CreatePenerimaan(&penerimaan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Data berhasil ditambahkan atau diperbarui"})
}

func (controller *PenerimaanController) UpdatePenerimaan(c echo.Context) error {
	kdTransaksiTerima := c.Param("kd_transaksi_terima")
	var penerimaan models.Penerimaan
	if err := c.Bind(&penerimaan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Input tidak valid"})
	}

	if err := models.UpdatePenerimaan(kdTransaksiTerima, &penerimaan); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, penerimaan)
}

func (controller *PenerimaanController) DeletePenerimaan(c echo.Context) error {
	kdTransaksiTerima := c.Param("kd_transaksi_terima")
	if err := models.DeletePenerimaan(kdTransaksiTerima); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Penerimaan berhasil dihapus"})
}

func (controller *PenerimaanController) DeleteAllPenerimaan(c echo.Context) error {
	nota := c.Param("no_terima")
	if err := models.DeletePenerimaanByNota(nota); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "All Transaksi successfully deleted"})
}
