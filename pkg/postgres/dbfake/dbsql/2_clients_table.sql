CREATE TABLE IF NOT EXISTS clients(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id VARCHAR NOT NULL,
    client_secret VARCHAR UNIQUE NOT NULL,
    user_id INTEGER NULL,
    managing_id VARCHAR (32),
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
