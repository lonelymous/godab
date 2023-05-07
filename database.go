package godab

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func OpenAndCreate(databaseConfig *DatabaseConfig, filename string) (err error) {
	err = Open(databaseConfig)
	if err != nil {
		return err
	}

	sqlFile, err := os.Open(filename)
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(sqlFile)

	fileScanner.Split(bufio.ScanLines)

	line := ""
	for fileScanner.Scan() {
		l := strings.TrimSpace(fileScanner.Text())

		if line == "" || strings.HasPrefix(l, "--") {
			continue
		}

		line = line + l

		if l[len(l)-1:] == ";" {
			Database.Exec(line)
			line = ""
		}
	}

	err = sqlFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// Close connection on database //defer it.
func CloseConnection() error {
	return Database.Close()
}
