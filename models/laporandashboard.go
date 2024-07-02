package models

import (
	"log"
)

type LaporanDashboard struct {
	TotBarang   string `json:"total_barang"`
	TotKategori string `json:"total_kategori"`
	TotPegawai  string `json:"total_pegawai"`
	TotPrefix   string `json:"total_prefix"`
	TotProvider string `json:"total_provider"`
	TotPulsa    string `json:"total_pulsa"`
	TotRole     string `json:"total_role"`
	TotSupplier string `json:"total_supplier"`
	TotUser     string `json:"total_user"`
}

func GetLaporanDashboard() ([]LaporanDashboard, error) {
	db := GetDB()
	rows, err := db.Query("SELECT * From laporan_dashboard")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanDashboard
	for rows.Next() {
		var b LaporanDashboard
		if err := rows.Scan(&b.TotBarang, &b.TotKategori, &b.TotPegawai, &b.TotPrefix, &b.TotProvider, &b.TotPulsa, &b.TotRole, &b.TotSupplier, &b.TotUser); err != nil {
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
