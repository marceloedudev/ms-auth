CREATE TABLE IF NOT EXISTS managing(
   id serial PRIMARY KEY,
   managing_id VARCHAR(32) NOT NULL,
   enabled BOOLEAN NOT NULL DEFAULT TRUE
);

INSERT INTO managing (managing_id) VALUES ('master');
