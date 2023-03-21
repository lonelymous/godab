package godab

import "database/sql"

// Exec
func Exec(queryString string, args ...any) (sql.Result, error) {
	return Database.Exec(queryString, args...)
}
