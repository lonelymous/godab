package godab

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

// Connection
// Open connection on database.
func OpenConnection(databaseConfig *DatabaseConfig) (err error) {
	database, err = sql.Open(databaseConfig.GetType(), databaseConfig.GetConnectionString())
	return err
}

// Close connection on database //defer it.
func CloseConnection() error {
	return database.Close()
}

// Basic
// Query gives back the all result
func Query(queryString string, args ...any) (*sql.Rows, error) {
	return database.Query(queryString, args...)
}

// QueryRow gives back only 1 row of result
func QueryRow(queryString string, args ...any) *sql.Row {
	return database.QueryRow(queryString, args...)
}

// Exec
func Exec(queryString string, args ...any) (sql.Result, error) {
	return database.Exec(queryString, args...)
}
