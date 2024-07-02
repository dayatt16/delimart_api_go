package models

import (
	"database/sql"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	KdPegawai   string `json:"kd_pegawai"`
	Username    string `json:"username"`
	NamaPegawai string `json:"nama_pegawai"`
	Password    string `json:"-"`
	KdRole      int    `json:"kd_role"`
	Live        int    `json:"-"`
}

// GetUserByUsername retrieves user data based on username.
func GetUserByUsername(username string) (*LoginUser, error) {
	db := GetDB()

	var user LoginUser
	err := db.QueryRow("SELECT u.kd_pegawai, u.username, u.password, p.kd_role, p.nama_pegawai, u.live FROM users u JOIN pegawai p ON u.kd_pegawai = p.kd_pegawai WHERE username = ?", username).Scan(&user.KdPegawai, &user.Username, &user.Password, &user.KdRole, &user.NamaPegawai, &user.Live)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		log.Printf("Error retrieving user: %v\n", err)
		return nil, err
	}

	return &user, nil
}

// CheckPassword compares the provided password with the stored password.
func CheckPassword(storedPassword, providedPassword string) bool {
	// If the stored password starts with "$2a$" or similar, assume it's hashed with bcrypt
	if strings.HasPrefix(storedPassword, "$2a$") || strings.HasPrefix(storedPassword, "$2b$") || strings.HasPrefix(storedPassword, "$2y$") {
		err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
		return err == nil
	}
	// Otherwise, assume it's a plain text password
	return storedPassword == providedPassword
}

// UpdateUserLiveStatus updates the live status of a user.
func UpdateUserLiveStatus(kdPegawai string, live int) error {
	db := GetDB()

	_, err := db.Exec("UPDATE users SET live = ? WHERE kd_pegawai = ?", live, kdPegawai)
	if err != nil {
		log.Printf("Error updating user live status for kd_pegawai %s: %v\n", kdPegawai, err)
		return err
	}

	log.Printf("User live status for kd_pegawai %s updated to %d", kdPegawai, live)
	return nil
}
