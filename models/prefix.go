package models

import (
	"database/sql"
	"errors"
	"log"
)

type Prefix struct {
	KodePrefix   string `json:"kd_prefix"`
	Prefix       string `json:"prefix"`
	KodeProvider string `json:"kd_provider"`
}

type VwPrefix struct {
	KodePrefix string `json:"kd_prefix"`
	Prefix     string `json:"prefix"`
	Provider   string `json:"provider"`
}

func GetAllPrefix() ([]VwPrefix, error) {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM vw_prefix")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var prefixes []VwPrefix
	for rows.Next() {
		var p VwPrefix
		if err := rows.Scan(&p.KodePrefix, &p.Prefix, &p.Provider); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		prefixes = append(prefixes, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v\n", err)
		return nil, err
	}

	return prefixes, nil
}

func GetProviderByPrefix(prefix string) (VwPrefix, error) {
	db := GetDB()
	var p VwPrefix

	row := db.QueryRow("SELECT * FROM vw_prefix WHERE prefix = ?", prefix)
	if err := row.Scan(&p.KodePrefix, &p.Prefix, &p.Provider); err != nil {
		if err == sql.ErrNoRows {
			return p, errors.New("p not found")
		}
		return p, err
	}
	return p, nil

}

func GetPrefixByKodePrefix(kdPrefix string) (Prefix, error) {
	db := GetDB()
	var prefix Prefix
	row := db.QueryRow("SELECT kd_prefix, prefix, kd_provider FROM prefix WHERE kd_prefix = ?", kdPrefix)
	if err := row.Scan(&prefix.KodePrefix, &prefix.Prefix, &prefix.KodeProvider); err != nil {
		if err == sql.ErrNoRows {
			return prefix, errors.New("Prefix not found")
		}
		return prefix, err
	}
	return prefix, nil
}

func CreatePrefix(prefix *Prefix) error {
	db := GetDB()
	sqlStatement := "INSERT INTO prefix (kd_prefix, prefix, kd_provider) VALUES (?, ?, ?)"
	_, err := db.Exec(sqlStatement, prefix.KodePrefix, prefix.Prefix, prefix.KodeProvider)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePrefixByKodePrefix(kdPrefix string, prefix *Prefix) (Prefix, error) {
	db := GetDB()
	sqlStatement := "UPDATE prefix SET prefix = ?, kd_provider = ? WHERE kd_prefix = ?"
	_, err := db.Exec(sqlStatement, prefix.Prefix, prefix.KodeProvider, kdPrefix)
	if err != nil {
		return *prefix, err
	}
	prefix.KodePrefix = kdPrefix
	return *prefix, nil
}

func DeletePrefixByKodePrefix(kdPrefix string) error {
	db := GetDB()
	sqlStatement := "DELETE FROM prefix WHERE kd_prefix = ?"
	_, err := db.Exec(sqlStatement, kdPrefix)
	if err != nil {
		return err
	}
	return nil
}
