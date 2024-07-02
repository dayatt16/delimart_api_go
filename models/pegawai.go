package models

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type Pegawai struct {
	KodePegawai  string `json:"kd_pegawai"`
	NamaPegawai  string `json:"nama_pegawai"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Alamat       string `json:"alamat"`
	Telepon      string `json:"telepon"`
	KodeRole     string `json:"kd_role"`
}

type VwPegawai struct {
	KodePegawai  string `json:"kd_pegawai"`
	NamaPegawai  string `json:"nama_pegawai"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Alamat       string `json:"alamat"`
	Telepon      string `json:"telepon"`
	Role     	 string `json:"role"`
}

func GetAllPegawai() ([]VwPegawai, error) {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM vw_pegawai")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pegawai []VwPegawai
	for rows.Next() {
		var p VwPegawai
		if err := rows.Scan(&p.KodePegawai, &p.NamaPegawai, &p.TanggalLahir, &p.JenisKelamin, &p.Alamat, &p.Telepon, &p.Role); err != nil {
			return nil, err
		}
		pegawai = append(pegawai, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pegawai, nil
}

func SearchPegawai(query string) ([]VwPegawai, error) {
	db := GetDB()
	query = "%" + strings.TrimSpace(query) + "%"
	rows, err := db.Query(`SELECT kd_pegawai, nama_pegawai, tanggal_lahir, jenis_kelamin, alamat, telepon, role FROM vw_pegawai WHERE nama_pegawai LIKE ? OR role LIKE ? OR kd_pegawai LIKE ?`, query, query, query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var hasil []VwPegawai
	for rows.Next() {
		var p VwPegawai
		if err := rows.Scan(&p.KodePegawai, &p.NamaPegawai, &p.TanggalLahir, &p.JenisKelamin, &p.Alamat, &p.Telepon, &p.Role); err != nil {
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


func GetPegawaiByKodePegawai(kdPegawai string) (Pegawai, error) {
	db := GetDB()
	var pegawai Pegawai
	row := db.QueryRow("SELECT kd_pegawai, nama_pegawai, tanggal_lahir, jenis_kelamin, alamat, telepon, kd_role FROM pegawai WHERE kd_pegawai = ?", kdPegawai)
	if err := row.Scan(&pegawai.KodePegawai, &pegawai.NamaPegawai, &pegawai.TanggalLahir, &pegawai.JenisKelamin, &pegawai.Alamat, &pegawai.Telepon, &pegawai.KodeRole); err != nil {
		if err == sql.ErrNoRows {
			return pegawai, errors.New("pegawai not found")
		}
		return pegawai, err
	}
	return pegawai, nil
}

func CreatePegawai(pegawai *Pegawai) error {
	db := GetDB()
	sqlStatement := "INSERT INTO pegawai (kd_pegawai, nama_pegawai, tanggal_lahir, jenis_kelamin, alamat, telepon, kd_role) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStatement, pegawai.KodePegawai, pegawai.NamaPegawai, pegawai.TanggalLahir, pegawai.JenisKelamin, pegawai.Alamat, pegawai.Telepon, pegawai.KodeRole)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePegawaiByKodePegawai(kdPegawai string, pegawai *Pegawai) error {
	db := GetDB()
	sqlStatement := "UPDATE pegawai SET nama_pegawai = ?, tanggal_lahir = ?, jenis_kelamin = ?, alamat = ?, telepon = ?, kd_role = ? WHERE kd_pegawai = ?"
	_, err := db.Exec(sqlStatement, pegawai.NamaPegawai, pegawai.TanggalLahir, pegawai.JenisKelamin, pegawai.Alamat, pegawai.Telepon, pegawai.KodeRole, kdPegawai)
	if err != nil {
		return err
	}
	return nil
}

func DeletePegawaiByKodePegawai(kdPegawai string) error {
	db := GetDB()
	sqlStatement := "DELETE FROM pegawai WHERE kd_pegawai = ?"
	_, err := db.Exec(sqlStatement, kdPegawai)
	if err != nil {
		return err
	}
	return nil
}
