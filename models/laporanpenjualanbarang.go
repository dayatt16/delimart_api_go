package models

import (
	"log"
)

type LaporanPenjualanBarang struct {
	NamaBarang string `json:"nama_barang"`
	Keuntungan string `json:"keuntungan"`
	Pendapatan string `json:"pendapatan"`
	QtyTerjual string `json:"qty_terjual"`
}

func GetLaporanPenjualanBarang() ([]LaporanPenjualanBarang, error) {
	db := GetDB()
	rows, err := db.Query("SELECT b.nama AS nama_barang, SUM(p.jumlah_beli) AS qty_terjual, SUM(p.jumlah_beli * b.harga_jual) AS pendapatan, ROUND(SUM(p.jumlah_beli * ((b.harga_jual - (b.harga_jual * b.diskon / 100)) - b.harga_beli))) AS keuntungan FROM penjualan p JOIN barang b ON p.kd_barang = b.kd_barang JOIN detail_penjualan dp ON p.struk = dp.struk WHERE DATE(p.created_at) = CURDATE() GROUP BY b.nama")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanPenjualanBarang
	for rows.Next() {
		var b LaporanPenjualanBarang
		if err := rows.Scan(&b.NamaBarang, &b.QtyTerjual, &b.Pendapatan, &b.Keuntungan); err != nil {
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
