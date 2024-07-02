package models

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type Kategori struct {
    KodeKategori string `json:"kd_kategori"`
    NamaKategori string `json:"nama_kategori"`
}

func GetAllKategori() ([]Kategori, error) {
    db := GetDB()
    rows, err := db.Query("SELECT kd_kategori, nama_kategori FROM kategori")
    if err != nil {
        log.Printf("Error executing query: %v\n", err)
        return nil, err
    }
    defer rows.Close()

    var kategoris []Kategori
    for rows.Next() {
        var k Kategori
        if err := rows.Scan(&k.KodeKategori, &k.NamaKategori); err != nil {
            log.Printf("Error scanning row: %v\n", err)
            return nil, err
        }
        kategoris = append(kategoris, k)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Error with rows: %v\n", err)
        return nil, err
    }

    return kategoris, nil
}

func SearchKategori(query string) ([]Kategori, error) {
	db := GetDB()
	query = "%" + strings.TrimSpace(query) + "%"
	rows, err := db.Query("SELECT kd_kategori, nama_kategori FROM kategori WHERE nama_kategori LIKE ? OR kd_kategori Like ?", query, query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var hasil []Kategori
	for rows.Next() {
		var k Kategori
		if err := rows.Scan(&k.KodeKategori, &k.NamaKategori); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		hasil = append(hasil, k)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v\n", err)
		return nil, err
	}

	return hasil, nil
}

func GetKategoriByKodeKategori(kdKategori string) (Kategori, error) {
    db := GetDB()
    var kategori Kategori
    row := db.QueryRow("SELECT kd_kategori, nama_kategori FROM kategori WHERE kd_kategori = ?", kdKategori)
    if err := row.Scan(&kategori.KodeKategori, &kategori.NamaKategori); err != nil {
        if err == sql.ErrNoRows {
            return kategori, errors.New("Kategori not found")
        }
        return kategori, err
    }
    return kategori, nil
}

func CreateKategori(kategori *Kategori) error {
    db := GetDB()
    sqlStatement := "INSERT INTO kategori (kd_kategori, nama_kategori) VALUES (?, ?)"
    _, err := db.Exec(sqlStatement, kategori.KodeKategori, kategori.NamaKategori)
    if err != nil {
        return err
    }
    return nil
}

func UpdateKategoriByKodeKategori(kdKategori string, kategori *Kategori) (Kategori, error) {
    db := GetDB()
    sqlStatement := "UPDATE kategori SET nama_kategori = ? WHERE kd_kategori = ?"
    _, err := db.Exec(sqlStatement, kategori.NamaKategori, kdKategori)
    if err != nil {
        return *kategori, err
    }
    kategori.KodeKategori = kdKategori
    return *kategori, nil
}

func DeleteKategoriByKodeKategori(kdKategori string) error {
    db := GetDB()
    sqlStatement := "DELETE FROM kategori WHERE kd_kategori = ?"
    _, err := db.Exec(sqlStatement, kdKategori)
    if err != nil {
        return err
    }
    return nil
}
