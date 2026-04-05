package db

import (
	"database/sql"
	_ "embed"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schemaSQL string

func Init() (*sql.DB, error) {
	// AppData\Local\LicenseManager\ 에 DB 저장
	appData := os.Getenv("LOCALAPPDATA")
	dbDir := filepath.Join(appData, "LION-FEEDER")
	os.MkdirAll(dbDir, 0755)

	dbPath := filepath.Join(dbDir, "licenses.db")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// 스키마 자동 적용
	if _, err := db.Exec(schemaSQL); err != nil {
		return nil, err
	}

	return db, nil
}