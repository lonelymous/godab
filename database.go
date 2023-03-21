package godab

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

// Connection
// Open connection on database.
func OpenConnection(databaseConfig *DatabaseConfig) (err error) {
	database, err = sql.Open(databaseConfig.GetDriver(), databaseConfig.GetConnectionString())
	return err
}

// Close connection on database //defer it.
func CloseConnection() error {
	return database.Close()
}
