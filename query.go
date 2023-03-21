package godab

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Query gives back the all result
func Query(queryString string, args ...any) (*sql.Rows, error) {
	return Database.Query(queryString, args...)
}

// QueryRow gives back only 1 row of result
func QueryRow(queryString string, args ...any) *sql.Row {
	return Database.QueryRow(queryString, args...)
}

// Query gives back the all result
func Queryx(queryString string, args ...any) (*sqlx.Rows, error) {
	return Database.Queryx(queryString, args...)
}

// QueryRow gives back only 1 row of result
func QueryRowx(queryString string, args ...any) *sqlx.Row {
	return Database.QueryRowx(queryString, args...)
}
