package models

import (
	"log"
	"time"
)

type DetailPenerimaan struct {
	Nota            string `json:"nota"`
	TanggalTerima   string `json:"tanggal_terima"`
	TotalItemTerima int    `json:"total_item_terima"`
	TotalQtyTerima  int    `json:"total_qty_terima"`
	TotalHargaBeli  int    `json:"total_harga_beli"`
	KdSupplier      string `json:"kd_supplier"`
	KdPegawai       string `json:"kd_pegawai"`
}

func CreateDetailPenerimaan(dp *DetailPenerimaan) error {
	db := GetDB()

	// Memastikan format tanggal sesuai dengan yang diharapkan oleh database
	t, err := time.Parse("2006-01-02 15:04:05", dp.TanggalTerima)
	if err != nil {
		return err
	}
	formattedDate := t.Format("2006-01-02 15:04:05")

	sqlStatement := `INSERT INTO detail_penerimaan (no_terima, tgl_terima, item_terima, qty_terima, total_harga_beli, kd_supplier, kd_pegawai) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(sqlStatement, dp.Nota, formattedDate, dp.TotalItemTerima, dp.TotalQtyTerima, dp.TotalHargaBeli, dp.KdSupplier, dp.KdPegawai)
	if err != nil {
		log.Printf("Error inserting into detail_penerimaan: %v", err)
		return err
	}
	return nil
}
