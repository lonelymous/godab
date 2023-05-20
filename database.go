package godab

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"

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

	Database, err := sqlx.Open(
		databaseConfig.GetDriver(),
		databaseConfig.GetConnectionStringWithoutDatabase(),
	)
	if err != nil {
		return nil, err
	}

	_, err = Database.Exec("USE " + databaseConfig.Name)
	if err != nil {
		log.Println("Database exists? " + strconv.FormatBool(!strings.Contains(err.Error(), "Unknown database")))
		if !strings.Contains(err.Error(), "Unknown database") {
			return nil, err
		}
		log.Println("Read SQL file..")
		sqlLines, err := ReadSQL(filename)
		if err != nil {
			return nil, err
		}

		log.Println("Execute SQL file.." + strconv.Itoa(len(sqlLines)) + " lines")

		for _, sql := range sqlLines {
			_, err := Database.Exec(sql)
			if err != nil {
				return nil, err
			}
			log.Println("Executed: " + sql)
		}
		return Database, nil
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
