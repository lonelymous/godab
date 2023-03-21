package godab

import (
	"github.com/jmoiron/sqlx"
)

var Database *sqlx.DB

// Connection
// Open connection on database.
func OpenConnection(databaseConfig *DatabaseConfig) (err error) {
	Database, err = sqlx.Connect(databaseConfig.GetDriver(), databaseConfig.GetConnectionString())
	return err
}

// Close connection on database //defer it.
func CloseConnection() error {
	return Database.Close()
}
