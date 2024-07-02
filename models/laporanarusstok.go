package models

import (
	"log"
)

type LaporanArusStok struct {
	NamaBarang   string `json:"nama_barang"`
	KodeBarang   string `json:"kd_barang"`
	BarangMasuk  string `json:"barang_masuk"`
	BarangKeluar string `json:"barang_keluar"`
}

func GetLaporanArusStok() ([]LaporanArusStok, error) {
	db := GetDB()
	rows, err := db.Query("SELECT b.nama AS nama_barang, b.kd_barang AS kd_barang, COALESCE(SUM(pn.jumlah_barang_terima), 0) AS barang_masuk, COALESCE(SUM(pj.jumlah_beli), 0) AS barang_keluar FROM barang b LEFT JOIN penerimaan pn ON pn.kd_barang = b.kd_barang AND DATE(pn.created_at) = CURDATE() AND EXISTS (SELECT 1 FROM detail_penerimaan dp WHERE dp.no_terima = pn.no_terima) LEFT JOIN penjualan pj ON pj.kd_barang = b.kd_barang AND DATE(pj.created_at) = CURDATE() AND EXISTS (SELECT 1 FROM detail_penjualan dj WHERE dj.struk = pj.struk) GROUP BY b.nama, b.kd_barang")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanArusStok
	for rows.Next() {
		var b LaporanArusStok
		if err := rows.Scan(&b.NamaBarang, &b.KodeBarang, &b.BarangMasuk, &b.BarangKeluar); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		laporan = append(laporan, b)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v\n", err)
		return nil, err
	}

	return laporan, nil
}
