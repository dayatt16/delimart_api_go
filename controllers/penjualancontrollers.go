package controllers

import (
	"Delimart/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PenjualanController struct{}

func (controller *PenjualanController) GetStrukAuto(c echo.Context) error {
	struk, err := models.GetStrukAuto()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// Mengembalikan respons JSON dengan urutan sebagai string
	return c.JSON(http.StatusOK, map[string]string{"struk": struk})
}

func (controller *PenjualanController) GetData(c echo.Context) error {
	struk := c.Param("struk")
	results, err := models.GetData(struk)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, results)
}

func (controller *PenjualanController) CreatePenjualan(c echo.Context) error {
	var penjualan models.Penjualan
	if err := c.Bind(&penjualan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Input tidak valid"})
	}

	// Periksa apakah stok valid untuk barang
	if penjualan.JenisProduk == "barang" && penjualan.JumlahBeli < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Minimal barang masuk 1"})
	}

	// Panggil fungsi CreateOrUpdatePenjualan dari models
	err := models.CreatePenjualan(&penjualan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Data berhasil ditambahkan atau diperbarui"})
}

func (controller *PenjualanController) UpdatePenjualan(c echo.Context) error {
	kdTransaksi := c.Param("kd_transaksi")

	// Parse JSON request body
	var penjualan models.Penjualan
	if err := c.Bind(&penjualan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
	}

	// Validate jumlah_beli
	if penjualan.JumlahBeli < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Minimal pembelian harus 1"})
	}

	// Update data in database
	err := models.UpdatePenjualan(kdTransaksi, &penjualan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Data berhasil diperbarui"})
}

func (controller *PenjualanController) DeletePenjualan(c echo.Context) error {
	kdtransaksi := c.Param("kd_transaksi")
	if err := models.DeletePenjualanByKodeTransaksi(kdtransaksi); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Transaksi successfully deleted"})
}

func (controller *PenjualanController) DeleteAllPenjualan(c echo.Context) error {
	struk := c.Param("struk")
	if err := models.DeletePenjualanByStruk(struk); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "All Transaksi successfully deleted"})
}
