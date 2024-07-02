package models

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type Pulsa struct {
	KodePulsa       string `json:"kd_pulsa"`
	NamaProdukPulsa string `json:"nama_produk_pulsa"`
	KodeProvider    string `json:"kd_provider"`
	Modal           string `json:"modal"`
	Harga           string `json:"harga"`
}

type VwPulsa struct {
	KodePulsa       string `json:"kd_pulsa"`
	NamaProdukPulsa string `json:"nama_produk_pulsa"`
	Provider        string `json:"provider"`
	Modal           string `json:"modal"`
	Harga           string `json:"harga"`
}

func GetAllPulsa() ([]VwPulsa, error) {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM vw_pulsa")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var hasil []VwPulsa
	for rows.Next() {
		var p VwPulsa
		if err := rows.Scan(&p.KodePulsa, &p.NamaProdukPulsa, &p.Provider, &p.Modal, &p.Harga); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		hasil = append(hasil, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v\n", err)
		return nil, err
	}

	return hasil, nil
}

func SearchPulsa(query string) ([]VwPulsa, error) {
	db := GetDB()
	query = "%" + strings.TrimSpace(query) + "%"
	rows, err := db.Query(`SELECT * FROM vw_pulsa WHERE kd_pulsa LIKE ? OR nama_produk_pulsa LIKE ? OR provider LIKE ?`, query, query, query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var hasil []VwPulsa
	for rows.Next() {
		var p VwPulsa
		if err := rows.Scan(&p.KodePulsa, &p.NamaProdukPulsa, &p.Provider, &p.Modal, &p.Harga); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		hasil = append(hasil, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v\n", err)
		return nil, err
	}

	return hasil, nil
}

func GetPulsaByProvider(provider string) ([]VwPulsa, error) {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM vw_pulsa WHERE provider = ?", provider)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var hasil []VwPulsa
	for rows.Next() {
		var p VwPulsa
		if err := rows.Scan(&p.KodePulsa, &p.NamaProdukPulsa, &p.Provider, &p.Modal, &p.Harga); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		hasil = append(hasil, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v\n", err)
		return nil, err
	}

	return hasil, nil
}

func GetPulsaByKodePulsa(kdPulsa string) (Pulsa, error) {
	db := GetDB()
	var p Pulsa
	row := db.QueryRow("SELECT kd_pulsa, nama_produk_pulsa, kd_provider, modal, harga FROM pulsa WHERE kd_pulsa = ?", kdPulsa)
	if err := row.Scan(&p.KodePulsa, &p.NamaProdukPulsa, &p.KodeProvider, &p.Modal, &p.Harga); err != nil {
		if err == sql.ErrNoRows {
			return p, errors.New("Pulsa not found")
		}
		return p, err
	}
	return p, nil
}

func CreatePulsa(pulsa *Pulsa) error {
	db := GetDB()
	sqlStatement := "INSERT INTO pulsa (kd_pulsa, nama_produk_pulsa, kd_provider, modal, harga) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStatement, pulsa.KodePulsa, pulsa.NamaProdukPulsa, pulsa.KodeProvider, pulsa.Modal, pulsa.Harga)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePulsaByKodePulsa(kdPulsa string, pulsa *Pulsa) (Pulsa, error) {
	db := GetDB()
	sqlStatement := "UPDATE pulsa SET nama_produk_pulsa = ?, kd_provider = ?, modal = ?, harga = ? WHERE kd_pulsa = ?"
	_, err := db.Exec(sqlStatement, pulsa.NamaProdukPulsa, pulsa.KodeProvider, pulsa.Modal, pulsa.Harga, kdPulsa)
	if err != nil {
		return *pulsa, err
	}
	pulsa.KodePulsa = kdPulsa
	return *pulsa, nil
}

func DeletePulsaByKodePulsa(kdPulsa string) error {
	db := GetDB()
	sqlStatement := "DELETE FROM pulsa WHERE kd_pulsa = ?"
	_, err := db.Exec(sqlStatement, kdPulsa)
	if err != nil {
		return err
	}
	return nil
}
