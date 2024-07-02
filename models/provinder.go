package models

import (
    "database/sql"
    "errors"
    "log"
)

type Provider struct {
    KodeProvider string `json:"kd_provider"`
    Provider     string `json:"provider"`
}

func GetAllProviders() ([]Provider, error) {
    db := GetDB()
    rows, err := db.Query("SELECT kd_provider, provider FROM provider")
    if err != nil {
        log.Printf("Error executing query: %v\n", err)
        return nil, err
    }
    defer rows.Close()

    var providers []Provider
    for rows.Next() {
        var p Provider
        if err := rows.Scan(&p.KodeProvider, &p.Provider); err != nil {
            log.Printf("Error scanning row: %v\n", err)
            return nil, err
        }
        providers = append(providers, p)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Error with rows: %v\n", err)
        return nil, err
    }

    return providers, nil
}

func GetProviderByKodeProvider(kdProvider string) (Provider, error) {
    db := GetDB()
    var provider Provider
    row := db.QueryRow("SELECT kd_provider, provider FROM provider WHERE kd_provider = ?", kdProvider)
    if err := row.Scan(&provider.KodeProvider, &provider.Provider); err != nil {
        if err == sql.ErrNoRows {
            return provider, errors.New("Provider not found")
        }
        return provider, err
    }
    return provider, nil
}

func CreateProvider(provider *Provider) error {
    db := GetDB()
    sqlStatement := "INSERT INTO provider (kd_provider, provider) VALUES (?, ?)"
    _, err := db.Exec(sqlStatement, provider.KodeProvider, provider.Provider)
    if err != nil {
        return err
    }
    return nil
}

func UpdateProviderByKodeProvider(kdProvider string, provider *Provider) (Provider, error) {
    db := GetDB()
    sqlStatement := "UPDATE provider SET provider = ? WHERE kd_provider = ?"
    _, err := db.Exec(sqlStatement, provider.Provider, kdProvider)
    if err != nil {
        return *provider, err
    }
    provider.KodeProvider = kdProvider
    return *provider, nil
}

func DeleteProviderByKodeProvider(kdProvider string) error {
    db := GetDB()
    sqlStatement := "DELETE FROM provider WHERE kd_provider = ?"
    _, err := db.Exec(sqlStatement, kdProvider)
    if err != nil {
        return err
    }
    return nil
}
