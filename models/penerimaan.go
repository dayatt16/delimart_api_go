package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Penerimaan struct {
	KodeTransaksiTerima  string `json:"kd_transaksi_terima"`
	NomorTerima          string `json:"no_terima"`
	KodeBarang           string `json:"kd_barang"`
	JumlahBarangDiterima int    `json:"jumlah_barang_terima"`
}
type GetPenerimaan struct {
	KodeTransaksiTerima  string `json:"kd_transaksi_terima"`
	NamaBarang           string `json:"nama_barang"`
	HargaBeli            int    `json:"harga_beli"`
	JumlahBarangDiterima int    `json:"jumlah_barang_terima"`
}

func GetNotaAuto() (string, error) {
	db := GetDB() // Fungsi GetDB() harus mengembalikan koneksi database yang sudah terbuka

	var maxNoTerima sql.NullString
	query := "SELECT MAX(no_terima) FROM penerimaan"
	err := db.QueryRow(query).Scan(&maxNoTerima)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return "", err
	}

	var urutan string
	if maxNoTerima.Valid {
		hitung := maxNoTerima.String[len(maxNoTerima.String)-3:]
		hitungInt := 0
		fmt.Sscanf(hitung, "%d", &hitungInt)
		hitungInt++
		urutan = fmt.Sprintf("T%s%03d", time.Now().Format("060102"), hitungInt)
	} else {
		urutan = fmt.Sprintf("T%s001", time.Now().Format("060102"))
	}

	return urutan, nil
}

func GetPenerimaanByNota(nota string) ([]GetPenerimaan, error) {
	db := GetDB()

	rows, err := db.Query("SELECT p.kd_transaksi_terima, b.nama, b.harga_beli, p.jumlah_barang_terima FROM penerimaan p JOIN barang b ON p.kd_barang = b.kd_barang WHERE p.no_terima = ? ORDER BY p.kd_transaksi_terima DESC", nota)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var penerimaans []GetPenerimaan

	for rows.Next() {
		var penerimaan GetPenerimaan
		err := rows.Scan(&penerimaan.KodeTransaksiTerima, &penerimaan.NamaBarang, &penerimaan.HargaBeli, &penerimaan.JumlahBarangDiterima)
		if err != nil {
			log.Printf("Error scanning penerimaan rows: %v\n", err)
			return nil, err
		}
		penerimaans = append(penerimaans, penerimaan)
	}

	// Check if no rows were returned
	if len(penerimaans) == 0 {
		log.Println("No data found for nota:", nota)
		// You may choose to return an empty slice or handle this case accordingly
		return []GetPenerimaan{}, nil
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over penerimaan rows: %v\n", err)
		return nil, err
	}

	return penerimaans, nil
}

// func GetPenerimaanByKodeTransaksiTerima(kdTransaksiTerima string) (*Penerimaan, error) {
// 	row := db.QueryRow("SELECT kd_transaksi_terima,  kd_barang, jumlah_barang_terima FROM penerimaan WHERE kd_transaksi_terima = ?", kdTransaksiTerima)

// 	var penerimaan Penerimaan
// 	if err := row.Scan(&penerimaan.KodeTransaksiTerima, &penerimaan.KodeBarang, &penerimaan.JumlahBarangDiterima); err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, errors.New("penerimaan tidak ditemukan")
// 		}
// 		return nil, err
// 	}

// 	return &penerimaan, nil
// }

func CreatePenerimaan(penerimaan *Penerimaan) error {
	db := GetDB()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var existingJumlah int
	query := "SELECT jumlah_barang_terima FROM penerimaan WHERE no_terima=? AND kd_barang=?"
	err = tx.QueryRow(query, penerimaan.NomorTerima, penerimaan.KodeBarang).Scan(&existingJumlah)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		// Barang belum ada, tambahkan data baru
		query = "INSERT INTO penerimaan (kd_transaksi_terima, no_terima, kd_barang, jumlah_barang_terima, created_at) VALUES (?, ?, ?, ?, ?)"
		_, err = tx.Exec(query, penerimaan.KodeTransaksiTerima, penerimaan.NomorTerima, penerimaan.KodeBarang, penerimaan.JumlahBarangDiterima, time.Now())
		if err != nil {
			return err
		}
	} else {
		// Barang sudah ada, update jumlahnya
		newJumlah := existingJumlah + penerimaan.JumlahBarangDiterima
		query = "UPDATE penerimaan SET jumlah_barang_terima=? WHERE no_terima=? AND kd_barang=?"
		_, err = tx.Exec(query, newJumlah, penerimaan.NomorTerima, penerimaan.KodeBarang)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// func UpdatePenerimaan(kdTransaksiTerima string, penerimaan *Penerimaan) (*Penerimaan, error) {
// 	_, err := db.Exec("UPDATE penerimaan SET  jumlah_barang_terima = ? WHERE kd_transaksi_terima = ?",
// 		penerimaan.JumlahBarangDiterima, kdTransaksiTerima)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func UpdatePenerimaan(kdTransaksiTerima string, penerimaan *Penerimaan) error {
	db := GetDB()
	sqlStatement := "UPDATE penerimaan SET  jumlah_barang_terima = ? WHERE kd_transaksi_terima = ?"
	_, err := db.Exec(sqlStatement, penerimaan.JumlahBarangDiterima, kdTransaksiTerima)
	if err != nil {
		return err
	}
	return nil
}

func DeletePenerimaan(kdTransaksiTerima string) error {
	_, err := db.Exec("DELETE FROM penerimaan WHERE kd_transaksi_terima = ?", kdTransaksiTerima)
	return err
}

func DeletePenerimaanByNota(nota string) error {
	sqlStatement := "DELETE FROM penerimaan WHERE no_terima = ?"
	_, err := db.Exec(sqlStatement, nota)
	if err != nil {
		return err
	}
	return nil
}
