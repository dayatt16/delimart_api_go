package models

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type Supplier struct {
	KodeSupplier string `json:"kd_supplier"`
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	Telepon      string `json:"telepon"`
}

func GetAllSuppliers() ([]Supplier, error) {
	db := GetDB()
	rows, err := db.Query("SELECT kd_supplier, nama, alamat, telepon FROM supplier")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []Supplier
	for rows.Next() {
		var s Supplier
		if err := rows.Scan(&s.KodeSupplier, &s.Nama, &s.Alamat, &s.Telepon); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, s)
	}

	if len(suppliers) == 0 {
        log.Println("No data found for supplier:")
        // You may choose to return an empty slice or handle this case accordingly
        return []Supplier{}, nil
    }


	return suppliers, nil
}

func SearchSupplier(query string) ([]Supplier, error) {
	db := GetDB()
	query = "%" + strings.TrimSpace(query) + "%"
	rows, err := db.Query("SELECT kd_supplier, nama, alamat, telepon FROM supplier WHERE nama LIKE ? OR alamat Like ? OR kd_supplier LIKE ?", query, query, query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var hasil []Supplier
	for rows.Next() {
		var s Supplier
		if err := rows.Scan(&s.KodeSupplier, &s.Nama, &s.Alamat, &s.Telepon); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		hasil = append(hasil, s)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v\n", err)
		return nil, err
	}

	return hasil, nil
}

func GetSupplierByKodeSupplier(kdSupplier string) (Supplier, error) {
	db := GetDB()
	var supplier Supplier
	row := db.QueryRow("SELECT kd_supplier, nama, alamat, telepon FROM supplier WHERE kd_supplier = ?", kdSupplier)
	if err := row.Scan(&supplier.KodeSupplier, &supplier.Nama, &supplier.Alamat, &supplier.Telepon); err != nil {
		if err == sql.ErrNoRows {
			return supplier, errors.New("Supplier not found")
		}
		return supplier, err
	}
	return supplier, nil
}
func GetKodeSupplierByNama(nama string) (Supplier, error) {
	db := GetDB()
	var supplier Supplier
	row := db.QueryRow("SELECT kd_supplier FROM supplier WHERE nama = ?", nama)
	if err := row.Scan(&supplier.KodeSupplier); err != nil {
		if err == sql.ErrNoRows {
			return supplier, errors.New("Supplier not found")
		}
		return supplier, err
	}
	
	return supplier, nil
}

func GetSupplierByNama(nama string) ([]Supplier, error) {
	db := GetDB()
	rows, err := db.Query("SELECT kd_supplier, nama, alamat, telepon FROM supplier WHERE nama = ?", nama)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []Supplier
	for rows.Next() {
		var s Supplier
		if err := rows.Scan(&s.KodeSupplier, &s.Nama, &s.Alamat, &s.Telepon); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, s)
	}

	if len(suppliers) == 0 {
        log.Println("No data found for supplier:", nama)
        // You may choose to return an empty slice or handle this case accordingly
        return []Supplier{}, nil
    }

	return suppliers, nil
}

func CreateSupplier(supplier *Supplier) error {
	db := GetDB()
	sqlStatement := "INSERT INTO supplier (kd_supplier, nama, alamat, telepon) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(sqlStatement, supplier.KodeSupplier, supplier.Nama, supplier.Alamat, supplier.Telepon)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSupplierByKodeSupplier(kdSupplier string, supplier *Supplier) (Supplier, error) {
	db := GetDB()
	sqlStatement := "UPDATE supplier SET nama = ?, alamat = ?, telepon = ? WHERE kd_supplier = ?"
	_, err := db.Exec(sqlStatement, supplier.Nama, supplier.Alamat, supplier.Telepon, kdSupplier)
	if err != nil {
		return *supplier, err
	}
	supplier.KodeSupplier = kdSupplier
	return *supplier, nil
}

func DeleteSupplierByKodeSupplier(kdSupplier string) error {
	db := GetDB()
	sqlStatement := "DELETE FROM supplier WHERE kd_supplier = ?"
	_, err := db.Exec(sqlStatement, kdSupplier)
	if err != nil {
		return err
	}
	return nil
}
