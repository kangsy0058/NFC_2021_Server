package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Mariadb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:hoseolab420@tcp(59.27.40.116:3306)/hoseo")
	if err != nil {
		return nil, err
	}

	return db, nil
}
