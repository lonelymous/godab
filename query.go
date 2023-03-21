package godab

import "database/sql"

// Query gives back the all result
func Query(queryString string, args ...any) (*sql.Rows, error) {
	return database.Query(queryString, args...)
}

// QueryRow gives back only 1 row of result
func QueryRow(queryString string, args ...any) *sql.Row {
	return database.QueryRow(queryString, args...)
}
