package model

type License struct {
	ID            int     `json:"id"`
	SoftwareID    int     `json:"software_id"`
	SoftwareName  string  `json:"software_name"`
	LicenseKey    string  `json:"license_key"`
	LicenseType   string  `json:"license_type"`
	SeatCount     int     `json:"seat_count"`
	UsedCount     int     `json:"used_count"`
	PurchaseDate  string  `json:"purchase_date"`
	ExpireDate    string  `json:"expire_date"`
	Cost          float64 `json:"cost"`
	VendorContact string  `json:"vendor_contact"`
	Notes         string  `json:"notes"`
}