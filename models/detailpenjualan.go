package models

import (
	"log"
	"time"
)

type DetailPenjualan struct {
	Struk       string `json:"struk"`
	TanggalJual string `json:"tanggal_jual"`
	TotalItem   int    `json:"total_item"`
	TotalQty    int    `json:"total_qty"`
	Diskon      int    `json:"diskon"`
	Pajak       int    `json:"pajak"`
	GrandTotal  int    `json:"grand_total"`
	Dibayar     int    `json:"dibayar"`
	Kembalian   int    `json:"kembalian"`
	KdPegawai   string `json:"kd_pegawai"`
}

func CreateDetailPenjualan(dpenj *DetailPenjualan) error {
	db := GetDB()
	t, err := time.Parse("2006-01-02 15:04:05", dpenj.TanggalJual)
	if err != nil {
		return err
	}
	formattedDate := t.Format("2006-01-02 15:04:05")
	sqlStatement := `INSERT INTO detail_penjualan (struk, tanggal_jual, total_item, total_qty, diskon, pajak, grand_total, dibayar, kembalian, kd_pegawai) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(sqlStatement, dpenj.Struk, formattedDate, dpenj.TotalItem, dpenj.TotalQty, dpenj.Diskon, dpenj.Pajak, dpenj.GrandTotal, dpenj.Dibayar, dpenj.Kembalian, dpenj.KdPegawai)
	if err != nil {
		log.Printf("Error inserting into detail_penerimaan: %v", err)
		return err
	}
	return nil
}
