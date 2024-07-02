package models

import (
	"log"
)

type LaporanPenjualanHome struct {
	Pendapatan     string `json:"total_pendapatan"`
	Keuntungan     string `json:"total_keuntungan"`
	Modal          string `json:"total_modal"`
	Transaksi      string `json:"total_transaksi"`
	QtyTerjual     string `json:"qty_terjual"`
	Rata2Penjualan string `json:"rata2_penjualan"`
	JmlhPegawai    string `json:"jumlah_pegawai"`
}

func GetLaporanPenjualanHome() ([]LaporanPenjualanHome, error) {
	db := GetDB()
	rows, err := db.Query("SELECT COALESCE(SUM(dp.grand_total - dp.pajak), 0) AS total_pendapatan_hari_ini, COALESCE(SUM(dp.grand_total - dp.pajak - dp.total_modal), 0) AS total_keuntungan_hari_ini, COALESCE(SUM(dp.total_modal), 0) AS total_modal_hari_ini, COALESCE(COUNT(DISTINCT dp.struk), 0) AS jumlah_transaksi_hari_ini, COALESCE(SUM(dp.jumlah_beli), 0) AS jumlah_qty_terjual_hari_ini, COALESCE(ROUND(AVG(rata2_penjualan)), 0) AS rata2_penjualan_hari_ini, COALESCE((SELECT COUNT(*) FROM pegawai), 0) AS jumlah_pegawai FROM ( SELECT dp.struk, dp.grand_total, dp.pajak, SUM(CASE WHEN p.jenis_produk = 'barang' THEN COALESCE(b.harga_beli * p.jumlah_beli, 0) WHEN p.jenis_produk = 'pulsa' THEN COALESCE(pl.modal * p.jumlah_beli, 0) ELSE 0 END) AS total_modal, SUM(p.jumlah_beli) AS jumlah_beli, (dp.grand_total - dp.pajak) AS rata2_penjualan FROM detail_penjualan dp JOIN penjualan p ON dp.struk = p.struk LEFT JOIN barang b ON p.kd_barang = b.kd_barang LEFT JOIN pulsa pl ON p.kd_pulsa = pl.kd_pulsa WHERE DATE(dp.tanggal_jual) = CURDATE() GROUP BY dp.struk, dp.grand_total, dp.pajak ) AS dp;")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanPenjualanHome
	for rows.Next() {
		var b LaporanPenjualanHome
		if err := rows.Scan(&b.Pendapatan, &b.Keuntungan, &b.Modal, &b.Transaksi, &b.QtyTerjual, &b.Rata2Penjualan, &b.JmlhPegawai); err != nil {
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
