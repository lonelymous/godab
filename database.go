package godab

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Database *sqlx.DB

// Connection
// Open connection on database.
func OpenConnection(databaseConfig *DatabaseConfig) (err error) {
	Database, err = sqlx.Connect(databaseConfig.GetDriver(), databaseConfig.GetConnectionString())
	return err
}

func Open(databaseConfig *DatabaseConfig) (err error) {
	Database, err = sqlx.Open(databaseConfig.GetDriver(), databaseConfig.GetConnectionStringWithoutDatabase())
	return err
}

// Close connection on database //defer it.
func CloseConnection() error {
	return Database.Close()
}
