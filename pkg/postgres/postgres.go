package postgres

import (
	"fmt"
	"log"
	"ms-auth/config"
	"ms-auth/pkg/postgres/dbfake"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

func InitPostgres(conf *config.Config) (*sqlx.DB, error) {

	source := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Postgres.Hostname, conf.Postgres.Port, conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.DBName)

	db, err := connectDB(conf.Postgres.DriverName, source)
	if err != nil {
		return nil, err
	}

	return db, nil

}

func connectDB(driverName string, dataSourceName string) (*sqlx.DB, error) {

	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, errors.New("database postgres connect")
	}

	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(60)
	db.SetConnMaxIdleTime(time.Second * 20)
	db.SetConnMaxLifetime(time.Second * 120)

	if err = db.Ping(); err != nil {
		return nil, errors.New("database postgres ping")
	}

	return db, nil

}

func InitPostgresTest() (*sqlx.DB, error) {

	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	err = dbfake.SqliteCreateTables(db)

	if err != nil {
		log.Fatalf("Create tables failed %v", err)
	}

	return db, nil

}
