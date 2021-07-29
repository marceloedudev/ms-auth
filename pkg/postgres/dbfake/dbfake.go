package dbfake

import (
	"fmt"
	"log"
	"ms-auth/pkg/utils"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
)

const (
	keyDBFakePath = "dbfake_path"
)

func GetEnvDBFakePath() string {
	path := os.Getenv(keyDBFakePath)

	os.Unsetenv(keyDBFakePath)

	return path
}

func SetEnvDBFakePath(path string) error {
	err := os.Setenv(keyDBFakePath, fmt.Sprintf("%s/pkg/postgres/dbfake/dbsql/*.sql", path))
	if err != nil {
		log.Fatalf("dbfake env failed: %v", err)
		return err
	}

	return nil
}

func SqliteCreateTables(db *sqlx.DB) (err error) {

	contents, err := utils.IOReadContentDir(GetEnvDBFakePath())

	if err != nil {
		return err
	}

	query := fmt.Sprintf("%s", strings.Join(contents, "\n"))

	_, err = db.Exec(query)

	if err != nil {
		return err
	}

	return nil

}
