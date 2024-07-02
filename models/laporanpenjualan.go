package models

import (
	"log"
)

type LaporanPenjualanHarian struct {
	Tanggal    string `json:"tanggal"`
	Keuntungan string `json:"keuntungan"`
	Pendapatan string `json:"pendapatan"`
	Jam        string `json:"jam"`
	Struk      string `json:"struk"`
}

type LaporanPenjualanBulanan struct {
	Tanggal        string `json:"tanggal"`
	Keuntungan     string `json:"keuntungan"`
	Pendapatan     string `json:"pendapatan"`
	TotalTransaksi string `json:"total_transaksi"`
}
type LaporanPenjualanTahunan struct {
	Bulan          string `json:"bulan"`
	Keuntungan     string `json:"keuntungan"`
	Pendapatan     string `json:"pendapatan"`
	TotalTransaksi string `json:"total_transaksi"`
}
type LaporanPenjualanAll struct {
	Tahun          string `json:"tahun"`
	Keuntungan     string `json:"keuntungan"`
	Pendapatan     string `json:"pendapatan"`
	TotalTransaksi string `json:"total_transaksi"`
}

func GetLaporanPenjualanHarian() ([]LaporanPenjualanHarian, error) {
	db := GetDB()
	rows, err := db.Query("SELECT DATE(dp.tanggal_jual) AS tanggal, TIME(dp.tanggal_jual) AS jam, (dp.grand_total - dp.pajak) AS pendapatan, (dp.grand_total - dp.pajak - SUM( CASE WHEN p.jenis_produk = 'barang' THEN (b.harga_beli * p.jumlah_beli) WHEN p.jenis_produk = 'pulsa' THEN (pl.modal * p.jumlah_beli) ELSE 0 END ) ) AS keuntungan, dp.struk FROM detail_penjualan dp JOIN penjualan p ON dp.struk = p.struk LEFT JOIN barang b ON p.kd_barang = b.kd_barang LEFT JOIN pulsa pl ON p.kd_pulsa = pl.kd_pulsa WHERE DATE(dp.tanggal_jual) = CURDATE() GROUP BY dp.struk")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanPenjualanHarian
	for rows.Next() {
		var b LaporanPenjualanHarian
		if err := rows.Scan(&b.Tanggal, &b.Jam, &b.Pendapatan, &b.Keuntungan, &b.Struk); err != nil {
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

func GetLaporanPenjualanBulanan() ([]LaporanPenjualanBulanan, error) {
	db := GetDB()
	rows, err := db.Query("SELECT DATE_FORMAT(dp.tanggal_jual, '%Y-%m-%d') AS tanggal, SUM(dp.grand_total - dp.pajak) AS total_pendapatan, SUM(dp.grand_total - dp.pajak - (SELECT SUM(CASE WHEN p.jenis_produk = 'barang' THEN (b.harga_beli * p.jumlah_beli) WHEN p.jenis_produk = 'pulsa' THEN (pl.modal * p.jumlah_beli) ELSE 0 END) FROM penjualan p LEFT JOIN barang b ON p.kd_barang = b.kd_barang LEFT JOIN pulsa pl ON p.kd_pulsa = pl.kd_pulsa WHERE p.struk = dp.struk GROUP BY p.struk) ) AS total_keuntungan, COUNT(dp.struk) AS total_transaksi FROM detail_penjualan dp WHERE MONTH(dp.tanggal_jual) = MONTH(CURDATE()) AND YEAR(dp.tanggal_jual) = YEAR(CURDATE()) GROUP BY DATE_FORMAT(dp.tanggal_jual, '%Y-%m-%d');")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanPenjualanBulanan
	for rows.Next() {
		var b LaporanPenjualanBulanan
		if err := rows.Scan(&b.Tanggal, &b.Pendapatan, &b.Keuntungan, &b.TotalTransaksi); err != nil {
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
func GetLaporanPenjualanTahunan() ([]LaporanPenjualanTahunan, error) {
	db := GetDB()
	rows, err := db.Query("SELECT DATE_FORMAT(dp.tanggal_jual, '%m') AS bulan, SUM(dp.grand_total - dp.pajak) AS total_pendapatan, SUM(dp.grand_total - dp.pajak - (SELECT SUM(CASE WHEN p.jenis_produk = 'barang' THEN (b.harga_beli * p.jumlah_beli) WHEN p.jenis_produk = 'pulsa' THEN (pl.modal * p.jumlah_beli) ELSE 0 END) FROM penjualan p LEFT JOIN barang b ON p.kd_barang = b.kd_barang LEFT JOIN pulsa pl ON p.kd_pulsa = pl.kd_pulsa WHERE p.struk = dp.struk GROUP BY p.struk) ) AS total_keuntungan, COUNT(dp.struk) AS total_transaksi FROM detail_penjualan dp WHERE YEAR(dp.tanggal_jual) = YEAR(CURDATE()) GROUP BY DATE_FORMAT(dp.tanggal_jual, '%m');")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanPenjualanTahunan
	for rows.Next() {
		var b LaporanPenjualanTahunan
		if err := rows.Scan(&b.Bulan, &b.Pendapatan, &b.Keuntungan, &b.TotalTransaksi); err != nil {
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
func GetLaporanPenjualanAll() ([]LaporanPenjualanAll, error) {
	db := GetDB()
	rows, err := db.Query("SELECT DATE_FORMAT(dp.tanggal_jual, '%Y') AS tahun, SUM(dp.grand_total - dp.pajak) AS total_pendapatan, SUM(dp.grand_total - dp.pajak - (SELECT SUM(CASE WHEN p.jenis_produk = 'barang' THEN (b.harga_beli * p.jumlah_beli) WHEN p.jenis_produk = 'pulsa' THEN (pl.modal * p.jumlah_beli) ELSE 0 END) FROM penjualan p LEFT JOIN barang b ON p.kd_barang = b.kd_barang LEFT JOIN pulsa pl ON p.kd_pulsa = pl.kd_pulsa WHERE p.struk = dp.struk GROUP BY p.struk) ) AS total_keuntungan, COUNT(dp.struk) AS total_transaksi FROM detail_penjualan dp GROUP BY DATE_FORMAT(dp.tanggal_jual, '%Y');")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanPenjualanAll
	for rows.Next() {
		var b LaporanPenjualanAll
		if err := rows.Scan(&b.Tahun, &b.Pendapatan, &b.Keuntungan, &b.TotalTransaksi); err != nil {
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
