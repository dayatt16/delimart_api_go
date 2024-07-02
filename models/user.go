package models

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type User struct {
	Id          string `json:"id"`
	KodePegawai string `json:"kd_pegawai"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Live        string `json:"live"`
}

type VwUser struct {
	Id       string `json:"id"`
	Pegawai  string `json:"nama_pegawai"`
	Username string `json:"username"`
	Live     string `json:"live"`
}

func GetAllUsers() ([]VwUser, error) {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM vw_user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []VwUser
	for rows.Next() {
		var u VwUser
		if err := rows.Scan(&u.Id, &u.Pegawai, &u.Username, &u.Live); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func SearchUser(query string) ([]VwUser, error) {
	db := GetDB()
	query = "%" + strings.TrimSpace(query) + "%"
	rows, err := db.Query("SELECT * FROM vw_user WHERE nama_pegawai LIKE ? OR username Like ?", query, query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var hasil []VwUser
	for rows.Next() {
		var u VwUser
		if err := rows.Scan(&u.Id, &u.Pegawai, &u.Username, &u.Live); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		hasil = append(hasil, u)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v\n", err)
		return nil, err
	}

	return hasil, nil
}

func GetUserByKodeUser(Id string) (User, error) {
	db := GetDB()
	var user User
	row := db.QueryRow("SELECT id, kd_pegawai, username, password, live FROM users WHERE id = ?", Id)
	if err := row.Scan(&user.Id, &user.KodePegawai, &user.Username, &user.Password, &user.Live); err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("User not found")
		}
		return user, err
	}
	return user, nil
}

func CreateUser(user *User) error {
	db := GetDB()
	sqlStatement := "INSERT INTO users (id, kd_pegawai, username, password, live) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStatement, user.Id, user.KodePegawai, user.Username, user.Password, user.Live)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserByKodeUser(Id string, user *User) (User, error) {
	db := GetDB()
	sqlStatement := "UPDATE users SET kd_pegawai = ?, username = ?, password = ?, live = ? WHERE id = ?"
	_, err := db.Exec(sqlStatement, user.KodePegawai, user.Username, user.Password, user.Live, Id)
	if err != nil {
		return *user, err
	}
	user.Id = Id
	return *user, nil
}

func DeleteUserByKodeUser(Id string) error {
	db := GetDB()
	sqlStatement := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(sqlStatement, Id)
	if err != nil {
		return err
	}
	return nil
}
