package models

import (
	"log"
)

type LaporanPenerimaanHarian struct {
	Tanggal     string `json:"tanggal"`
	Pengeluaran string `json:"pengeluaran"`
	Supplier    string `json:"supplier"`
	Jam         string `json:"jam"`
	Nota        string `json:"nota"`
}

type LaporanPenerimaanBulanan struct {
	Tanggal        string `json:"tanggal"`
	Pengeluaran    string `json:"pengeluaran"`
	TotalTransaksi string `json:"total_transaksi"`
}

type LaporanPenerimaanTahunan struct {
	Bulanan        string `json:"bulan"`
	Pengeluaran    string `json:"pengeluaran"`
	TotalTransaksi string `json:"total_transaksi"`
}

type LaporanPenerimaanAll struct {
	Tahunan        string `json:"Tahunan"`
	Pengeluaran    string `json:"pengeluaran"`
	TotalTransaksi string `json:"total_transaksi"`
}

func GetLaporanPenerimaanHarian() ([]LaporanPenerimaanHarian, error) {
	db := GetDB()
	rows, err := db.Query("SELECT DATE_FORMAT(dp.tgl_terima, '%Y-%m-%d') AS tanggal, dp.total_harga_beli AS pengeluaran, s.nama AS supplier, TIME(dp.tgl_terima) AS jam, dp.no_terima FROM detail_penerimaan dp JOIN supplier s ON dp.kd_supplier = s.kd_supplier WHERE DATE(dp.tgl_terima) = CURDATE();")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanPenerimaanHarian
	for rows.Next() {
		var b LaporanPenerimaanHarian
		if err := rows.Scan(&b.Tanggal, &b.Pengeluaran, &b.Supplier, &b.Jam, &b.Nota); err != nil {
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

func GetLaporanPenerimaanBulanan() ([]LaporanPenerimaanBulanan, error) {
	db := GetDB()
	rows, err := db.Query("SELECT DATE(tgl_terima) AS tanggal_terima, SUM(total_harga_beli) AS total_pengeluaran, COUNT(*) AS total_transaksi FROM detail_penerimaan dp JOIN penerimaan p ON dp.no_terima = p.no_terima WHERE MONTH(tgl_terima) = MONTH(CURDATE()) GROUP BY DATE(tgl_terima) ORDER BY tanggal_terima ASC")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanPenerimaanBulanan
	for rows.Next() {
		var b LaporanPenerimaanBulanan
		if err := rows.Scan(&b.Tanggal, &b.Pengeluaran, &b.TotalTransaksi); err != nil {
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

func GetLaporanPenerimaanTahunan() ([]LaporanPenerimaanTahunan, error) {
	db := GetDB()
	rows, err := db.Query("SELECT DATE_FORMAT(tgl_terima,'%m') AS bulan, SUM(total_harga_beli) AS total_pengeluaran, COUNT(*) AS total_transaksi FROM detail_penerimaan dp JOIN penerimaan p ON dp.no_terima = p.no_terima WHERE YEAR(tgl_terima) = YEAR(CURDATE()) GROUP BY DATE_FORMAT(tgl_terima,'%m') ORDER BY bulan ASC")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanPenerimaanTahunan
	for rows.Next() {
		var b LaporanPenerimaanTahunan
		if err := rows.Scan(&b.Bulanan, &b.Pengeluaran, &b.TotalTransaksi); err != nil {
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

func GetLaporanPenerimaanAll() ([]LaporanPenerimaanAll, error) {
	db := GetDB()
	rows, err := db.Query("SELECT DATE_FORMAT(tgl_terima,'%Y') AS tahun, SUM(total_harga_beli) AS total_pengeluaran, COUNT(*) AS total_transaksi FROM detail_penerimaan dp JOIN penerimaan p ON dp.no_terima = p.no_terima GROUP BY DATE_FORMAT(tgl_terima,'%Y') ORDER BY tahun ASC")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanPenerimaanAll
	for rows.Next() {
		var b LaporanPenerimaanAll
		if err := rows.Scan(&b.Tahunan, &b.Pengeluaran, &b.TotalTransaksi); err != nil {
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
