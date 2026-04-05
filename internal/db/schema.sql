
CREATE TABLE IF NOT EXISTS software (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT    NOT NULL,
    vendor      TEXT,
    category    TEXT,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS licenses (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    software_id     INTEGER REFERENCES software(id),
    license_key     TEXT,
    license_type    TEXT,
    seat_count      INTEGER DEFAULT 1,
    used_count      INTEGER DEFAULT 0,
    purchase_date   DATE,
    expire_date     DATE,
    cost            REAL,
    vendor_contact  TEXT,
    notes           TEXT,
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS assignments (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    license_id  INTEGER REFERENCES licenses(id),
    user_name   TEXT,
    department  TEXT,
    device_name TEXT,
    assigned_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    returned_at DATETIME
);