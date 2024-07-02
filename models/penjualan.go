package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

type NullString struct {
	sql.NullString
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		ns.Valid = true
		ns.String = *s
	} else {
		ns.Valid = false
	}
	return nil
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

type Transaksi struct {
	KdTransaksi string  `json:"kd_transaksi"`
	Nama        string  `json:"nama"`
	HargaJual   int     `json:"harga_jual"`
	JumlahBeli  int     `json:"jumlah_beli"`
	Diskon      float64 `json:"diskon"`
}

type Penjualan struct {
	KodeTransaksi string     `json:"kd_transaksi"`
	Struk         string     `json:"struk"`
	KodeBarang    NullString `json:"kd_barang"`
	JumlahBeli    int        `json:"jumlah_beli"`
	KodePulsa     string     `json:"kd_pulsa"`
	JenisProduk   string     `json:"jenis_produk"`
}

func GetStrukAuto() (string, error) {
	db := GetDB() // Fungsi GetDB() harus mengembalikan koneksi database yang sudah terbuka

	var struk sql.NullString
	query := "SELECT MAX(struk) FROM penjualan"
	err := db.QueryRow(query).Scan(&struk)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return "", err
	}

	var urutan string
	if struk.Valid {
		hitung := struk.String[len(struk.String)-3:]
		hitungInt := 0
		fmt.Sscanf(hitung, "%d", &hitungInt)
		hitungInt++
		urutan = fmt.Sprintf("J%s%03d", time.Now().Format("060102"), hitungInt)
	} else {
		urutan = fmt.Sprintf("J%s001", time.Now().Format("060102"))
	}

	return urutan, nil
}

func GetData(struk string) ([]Transaksi, error) {
	db := GetDB() // Fungsi GetDB() harus mengembalikan koneksi database yang sudah terbuka
	var results []Transaksi

	// Query untuk mendapatkan data barang
	queryBarang := `
		SELECT penjualan.kd_transaksi, barang.nama, barang.harga_jual, barang.diskon, penjualan.jumlah_beli 
		FROM penjualan 
		JOIN barang ON barang.kd_barang = penjualan.kd_barang 
		WHERE penjualan.struk = ? AND penjualan.jenis_produk = 'barang' 
		ORDER BY penjualan.kd_transaksi DESC
	`

	rowsBarang, err := db.Query(queryBarang, struk)
	if err != nil {
		log.Printf("Error executing queryBarang: %v\n", err)
		return results, err
	}
	defer rowsBarang.Close()

	for rowsBarang.Next() {
		var transaksi Transaksi
		if err := rowsBarang.Scan(&transaksi.KdTransaksi, &transaksi.Nama, &transaksi.HargaJual, &transaksi.Diskon, &transaksi.JumlahBeli); err != nil {
			log.Printf("Error scanning rowBarang: %v\n", err)
			return results, err
		}
		results = append(results, transaksi)
	}

	// Query untuk mendapatkan data pulsa
	queryPulsa := `
		SELECT penjualan.kd_transaksi, pulsa.nama_produk_pulsa AS nama, pulsa.harga AS harga_jual, 0 AS diskon, penjualan.jumlah_beli 
		FROM penjualan 
		JOIN pulsa ON pulsa.kd_pulsa = penjualan.kd_pulsa 
		WHERE penjualan.struk = ? AND penjualan.jenis_produk = 'pulsa' 
		ORDER BY penjualan.kd_transaksi DESC
	`

	rowsPulsa, err := db.Query(queryPulsa, struk)
	if err != nil {
		log.Printf("Error executing queryPulsa: %v\n", err)
		return results, err
	}
	defer rowsPulsa.Close()

	for rowsPulsa.Next() {
		var transaksi Transaksi
		if err := rowsPulsa.Scan(&transaksi.KdTransaksi, &transaksi.Nama, &transaksi.HargaJual, &transaksi.Diskon, &transaksi.JumlahBeli); err != nil {
			log.Printf("Error scanning rowPulsa: %v\n", err)
			return results, err
		}
		results = append(results, transaksi)
	}

	// Check if no data was found for both queries
	if len(results) == 0 {
		log.Println("No data found for nota:", struk)
		// You may choose to return an empty slice or handle this case accordingly
		return []Transaksi{}, nil
	}

	return results, nil
}

func CreatePenjualan(penjualan *Penjualan) error {
	db := GetDB()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if penjualan.JenisProduk == "barang" {
		// Cek stok barang
		var stok int
		err = tx.QueryRow("SELECT stok FROM barang WHERE kd_barang = ?", penjualan.KodeBarang).Scan(&stok)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("kode Barang Tidak ada")
			}
			return err
		}

		if stok < 1 {
			return errors.New("maaf stok habis")
		}

		// Cek apakah barang sudah ada dalam transaksi penjualan
		var existingJumlah int
		query := "SELECT jumlah_beli FROM penjualan WHERE struk=? AND kd_barang=?"
		err = tx.QueryRow(query, penjualan.Struk, penjualan.KodeBarang).Scan(&existingJumlah)

		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if err == sql.ErrNoRows {
			// Barang belum ada, tambahkan data baru
			query = "INSERT INTO penjualan (kd_transaksi, struk, kd_barang, jumlah_beli,  jenis_produk,created_at) VALUES (?, ?, ?, ?, ?, ?)"
			_, err = tx.Exec(query, penjualan.KodeTransaksi, penjualan.Struk, penjualan.KodeBarang, penjualan.JumlahBeli, penjualan.JenisProduk, time.Now())
			if err != nil {
				return err
			}
		} else {
			// Barang sudah ada, update jumlahnya
			newJumlah := existingJumlah + penjualan.JumlahBeli
			query = "UPDATE penjualan SET jumlah_beli=? WHERE struk=? AND kd_barang=?"
			_, err = tx.Exec(query, newJumlah, penjualan.Struk, penjualan.KodeBarang)
			if err != nil {
				return err
			}
		}
	} else if penjualan.JenisProduk == "pulsa" {
		// Cek apakah pulsa sudah ada dalam transaksi penjualan
		var existingJumlah int
		query := "SELECT jumlah_beli FROM penjualan WHERE struk=? AND kd_pulsa=?"
		err = tx.QueryRow(query, penjualan.Struk, penjualan.KodePulsa).Scan(&existingJumlah)

		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if err == sql.ErrNoRows {
			// Pulsa belum ada, tambahkan data baru
			query = "INSERT INTO penjualan (kd_transaksi, struk, kd_pulsa, jumlah_beli,  jenis_produk,created_at) VALUES (?, ?, ?, ?, ?, ?)"
			_, err = tx.Exec(query, penjualan.KodeTransaksi, penjualan.Struk, penjualan.KodePulsa, penjualan.JumlahBeli, penjualan.JenisProduk, time.Now())
			if err != nil {
				return err
			}
		} else {
			// Pulsa sudah ada, update jumlahnya
			newJumlah := existingJumlah + penjualan.JumlahBeli
			query = "UPDATE penjualan SET jumlah_beli=? WHERE struk=? AND kd_pulsa=?"
			_, err = tx.Exec(query, newJumlah, penjualan.Struk, penjualan.KodePulsa)
			if err != nil {
				return err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func UpdatePenjualan(kdTransaksi string, penjualan *Penjualan) error {
	db := GetDB()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Cek data penjualan berdasarkan kd_transaksi
	var existingPenjualan Penjualan
	query := "SELECT jumlah_beli, jenis_produk, kd_barang FROM penjualan WHERE kd_transaksi = ?"
	err = tx.QueryRow(query, kdTransaksi).Scan(&existingPenjualan.JumlahBeli, &existingPenjualan.JenisProduk, &existingPenjualan.KodeBarang)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("data penjualan tidak ditemukan")
		}
		return err
	}

	// Proses update sesuai jenis produk
	if existingPenjualan.JenisProduk == "barang" {
		// Cek stok barang
		var stok int
		err = tx.QueryRow("SELECT stok FROM barang WHERE kd_barang = ?", existingPenjualan.KodeBarang.String).Scan(&stok)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("kode barang tidak ada")
			}
			return err
		}

		// Hitung stok terbaru
		stokTerbaru := stok + existingPenjualan.JumlahBeli
		if penjualan.JumlahBeli > stokTerbaru {
			return errors.New("stok barang tersisa " + fmt.Sprint(stok))
		}

		// Update jumlah beli
		query = "UPDATE penjualan SET jumlah_beli = ? WHERE kd_transaksi = ?"
		_, err = tx.Exec(query, penjualan.JumlahBeli, kdTransaksi)
		if err != nil {
			return err
		}

		// Kurangi stok barang
		newStok := stokTerbaru - penjualan.JumlahBeli
		query = "UPDATE barang SET stok = ? WHERE kd_barang = ?"
		_, err = tx.Exec(query, newStok, existingPenjualan.KodeBarang.String)
		if err != nil {
			return err
		}

	} else if existingPenjualan.JenisProduk == "pulsa" {
		// Update jumlah beli pulsa
		query = "UPDATE penjualan SET jumlah_beli = ? WHERE kd_transaksi = ?"
		_, err = tx.Exec(query, penjualan.JumlahBeli, kdTransaksi)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
func DeletePenjualanByKodeTransaksi(kdtransaksi string) error {
	db := GetDB()
	sqlStatement := "DELETE FROM penjualan WHERE kd_transaksi = ?"
	_, err := db.Exec(sqlStatement, kdtransaksi)
	if err != nil {
		return err
	}
	return nil
}
func DeletePenjualanByStruk(struk string) error {
	db := GetDB()
	sqlStatement := "DELETE FROM penjualan WHERE struk = ?"
	_, err := db.Exec(sqlStatement, struk)
	if err != nil {
		return err
	}
	return nil
}
