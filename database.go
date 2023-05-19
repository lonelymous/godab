package godab

import (
	"bufio"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Database *sqlx.DB

// Connection
// Connect to database.
func Connect(databaseConfig *DatabaseConfig) (*sqlx.DB, error) {
	var err error
	Database, err = sqlx.Connect(databaseConfig.GetDriver(), databaseConfig.GetConnectionString())
	return Database, err
}

// Open connection on database.
func Open(databaseConfig *DatabaseConfig) (*sqlx.DB, error) {
	var err error
	Database, err = sqlx.Open(databaseConfig.GetDriver(), databaseConfig.GetConnectionStringWithoutDatabase())
	return Database, err
}

// Open and create database from sql file.
func OpenAndCreate(databaseConfig *DatabaseConfig, filename string) (*sqlx.DB, error) {
	var err error

	Database, err = sqlx.Open(
		databaseConfig.GetDriver(),
		databaseConfig.GetConnectionStringWithoutDatabase(),
	)
	if err != nil {
		return nil, err
	}

	sqlLines, err := ReadSQL(filename)
	if err != nil {
		return nil, err
	}

	for _, sql := range sqlLines {
		_, err := Database.Exec(sql)
		if err != nil {
			return nil, err
		}
	}

	_, err = Database.Exec("USE " + databaseConfig.Name)
	if err != nil {
		return nil, err
	}

	return Database, err
}

// Close connection on database //defer it.
func CloseConnection() error {
	return Database.Close()
}

// Read SQL file and return array of SQL query.
func ReadSQL(filename string) ([]string, error) {

	sqlFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(sqlFile)

	fileScanner.Split(bufio.ScanLines)

	lines := []string{}
	line := ""
	for fileScanner.Scan() {
		l := strings.TrimSpace(fileScanner.Text())

		if l == "" || strings.HasPrefix(l, "--") {
			continue
		}

		line = line + l

		if l[len(l)-1:] == ";" {
			lines = append(lines, line)
			line = ""
		}
	}

	return lines, sqlFile.Close()
}
