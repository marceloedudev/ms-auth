CREATE TABLE IF NOT EXISTS managing(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    managing_id VARCHAR(32) NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT TRUE
);
