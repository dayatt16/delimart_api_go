package models

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type Barang struct {
	KodeBarang   string `json:"kd_barang"`
	KodeSupplier string `json:"kd_supplier"`
	Nama         string `json:"nama"`
	KodeKategori string `json:"kd_kategori"`
	HargaBeli    string    `json:"harga_beli"`
	HargaJual    string    `json:"harga_jual"`
	Diskon       string    `json:"diskon"`
	Stok         string    `json:"stok"`
}

type VwBarang struct {
	KodeBarang string `json:"kd_barang"`
	Supplier   string `json:"supplier"`
	Nama       string `json:"nama"`
	Kategori   string `json:"kategori"`
	HargaBeli  int    `json:"harga_beli"`
	HargaJual  int    `json:"harga_jual"`
	Diskon     int    `json:"diskon"`
	Stok       int    `json:"stok"`
}

func GetAllBarang() ([]VwBarang, error) {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM vw_barang")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var barangs []VwBarang
	for rows.Next() {
		var b VwBarang
		if err := rows.Scan(&b.KodeBarang, &b.Supplier, &b.Nama, &b.Kategori, &b.HargaBeli, &b.HargaJual, &b.Diskon, &b.Stok); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		barangs = append(barangs, b)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v\n", err)
		return nil, err
	}

	return barangs, nil
}

func SearchBarang(query string) ([]VwBarang, error) {
	db := GetDB()
	query = "%" + strings.TrimSpace(query) + "%"
	rows, err := db.Query(`SELECT * FROM vw_barang WHERE kd_barang LIKE ? OR nama LIKE ? OR kategori LIKE ? OR supplier Like ?`, query, query, query, query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var hasil []VwBarang
	for rows.Next() {
		var b VwBarang
		if err := rows.Scan(&b.KodeBarang, &b.Supplier, &b.Nama, &b.Kategori, &b.HargaJual, &b.HargaBeli,  &b.Diskon, &b.Stok); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		hasil = append(hasil, b)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v\n", err)
		return nil, err
	}

	return hasil, nil
}

func GetBarangByKodeBarang(kdBarang string) (Barang, error) {
	db := GetDB()
	var barang Barang
	row := db.QueryRow("SELECT kd_barang, kd_supplier, nama, kd_kategori, harga_beli, harga_jual, diskon, stok FROM barang WHERE kd_barang = ?", kdBarang)
	if err := row.Scan(&barang.KodeBarang, &barang.KodeSupplier, &barang.Nama, &barang.KodeKategori, &barang.HargaBeli, &barang.HargaJual, &barang.Diskon, &barang.Stok); err != nil {
		if err == sql.ErrNoRows {
			return barang, errors.New("barang not found")
		}
		return barang, err
	}
	return barang, nil
}

func GetBarangByKodeSupp(kdSupplier string) ([]Barang, error) {
	db := GetDB()
	var barangs []Barang

	rows, err := db.Query("SELECT  nama FROM barang WHERE kd_supplier = ?", kdSupplier)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var barang Barang
		err := rows.Scan(&barang.Nama)
		if err != nil {
			return nil, err
		}
		barangs = append(barangs, barang)
	}
	// Check if no rows were returned
	if len(barangs) == 0 {
		log.Println("No data found for nota:", kdSupplier)
		// You may choose to return an empty slice or handle this case accordingly
		return []Barang{}, nil
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return barangs, nil
}

func GetBarangByNama(nama string) (Barang, error) {
	db := GetDB()
	var barang Barang

	row := db.QueryRow("SELECT kd_barang, harga_beli FROM barang WHERE nama = ?", nama)
	if err := row.Scan(&barang.KodeBarang, &barang.HargaBeli); err != nil {
		if err == sql.ErrNoRows {
			return barang, errors.New("barang not found")
		}
		return barang, err
	}
	return barang, nil
}

func CreateBarang(barang *Barang) error {
	db := GetDB()
	sqlStatement := "INSERT INTO barang (kd_barang, kd_supplier, nama, kd_kategori, harga_beli, harga_jual, diskon, stok) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStatement, barang.KodeBarang, barang.KodeSupplier, barang.Nama, barang.KodeKategori, barang.HargaBeli, barang.HargaJual, barang.Diskon, barang.Stok)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBarangByKodeBarang(kdBarang string, barang *Barang) (Barang, error) {
	db := GetDB()
	sqlStatement := "UPDATE barang SET kd_supplier = ?, nama = ?, kd_kategori = ?, harga_beli = ?, harga_jual = ?, diskon = ?, stok = ? WHERE kd_barang = ?"
	_, err := db.Exec(sqlStatement, barang.KodeSupplier, barang.Nama, barang.KodeKategori, barang.HargaBeli, barang.HargaJual, barang.Diskon, barang.Stok, kdBarang)
	if err != nil {
		return *barang, err
	}
	barang.KodeBarang = kdBarang
	return *barang, nil
}

func DeleteBarangByKodeBarang(kdBarang string) error {
	db := GetDB()
	sqlStatement := "DELETE FROM barang WHERE kd_barang = ?"
	_, err := db.Exec(sqlStatement, kdBarang)
	if err != nil {
		return err
	}
	return nil
}
