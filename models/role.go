package models

import (
    "database/sql"
    "errors"
    "log"
)

type Role struct {
    KodeRole string `json:"kd_role"`
    Role     string `json:"role"`
}

func GetAllRoles() ([]Role, error) {
    db := GetDB()
    rows, err := db.Query("SELECT kd_role, role FROM role")
    if err != nil {
        log.Printf("Error executing query: %v\n", err)
        return nil, err
    }
    defer rows.Close()

    var roles []Role
    for rows.Next() {
        var r Role
        if err := rows.Scan(&r.KodeRole, &r.Role); err != nil {
            log.Printf("Error scanning row: %v\n", err)
            return nil, err
        }
        roles = append(roles, r)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Error with rows: %v\n", err)
        return nil, err
    }

    return roles, nil
}

func GetRoleByKodeRole(kdRole string) (Role, error) {
    db := GetDB()
    var role Role
    row := db.QueryRow("SELECT kd_role, role FROM role WHERE kd_role = ?", kdRole)
    if err := row.Scan(&role.KodeRole, &role.Role); err != nil {
        if err == sql.ErrNoRows {
            return role, errors.New("Role not found")
        }
        return role, err
    }
    return role, nil
}

func CreateRole(role *Role) error {
    db := GetDB()
    sqlStatement := "INSERT INTO role (kd_role, role) VALUES (?, ?)"
    _, err := db.Exec(sqlStatement, role.KodeRole, role.Role)
    if err != nil {
        return err
    }
    return nil
}

func UpdateRoleByKodeRole(kdRole string, role *Role) (Role, error) {
    db := GetDB()
    sqlStatement := "UPDATE role SET role = ? WHERE kd_role = ?"
    _, err := db.Exec(sqlStatement, role.Role, kdRole)
    if err != nil {
        return *role, err
    }
    role.KodeRole = kdRole
    return *role, nil
}

func DeleteRoleByKodeRole(kdRole string) error {
    db := GetDB()
    sqlStatement := "DELETE FROM role WHERE kd_role = ?"
    _, err := db.Exec(sqlStatement, kdRole)
    if err != nil {
        return err
    }
    return nil
}
