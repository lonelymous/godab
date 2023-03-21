package godab

import "database/sql"

// Exec
func Exec(queryString string, args ...any) (sql.Result, error) {
	return database.Exec(queryString, args...)
}
