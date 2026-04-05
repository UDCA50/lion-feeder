package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func ListLicenses(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT l.id, l.software_id, s.name, l.license_type,
			       l.seat_count, l.used_count, l.expire_date
			FROM licenses l
			LEFT JOIN software s ON s.id = l.software_id
			ORDER BY l.expire_date ASC
		`)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()

		var results []map[string]any
		for rows.Next() {
			var id, softwareID, seatCount, usedCount int
			var name, licenseType, expireDate string
			rows.Scan(&id, &softwareID, &name, &licenseType,
				&seatCount, &usedCount, &expireDate)
			results = append(results, map[string]any{
				"id":           id,
				"software_id":  softwareID,
				"name":         name,
				"license_type": licenseType,
				"seat_count":   seatCount,
				"used_count":   usedCount,
				"expire_date":  expireDate,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}

func CreateLicense(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Phase 2에서 구현
		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateLicense(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteLicense(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
}