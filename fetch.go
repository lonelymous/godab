package godab

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func FetchRowsToAny[T any](rows *sql.Rows, result *[]T) error {
	for rows.Next() {
		var row T

		err := rows.Scan(&row)
		if err != nil {
			return err
		}

		*result = append(*result, row)
	}

	return nil
}

func FetchRowToAny[T any](rows *sql.Row, result *T) error {
	err := rows.Scan(&result)
	if err != nil {
		return err
	}

	return nil
}

func FetchXRowsToStruct[T any](rows *sqlx.Rows, result *[]T) error {
	for rows.Next() {
		var row T

		err := rows.StructScan(&row)
		if err != nil {
			return err
		}

		*result = append(*result, row)
	}

	return nil
}

func FetchXRowToStruct[T any](rows *sqlx.Row, result *T) error {
	err := rows.StructScan(&result)
	if err != nil {
		return err
	}

	return nil
}

func FetchXRowsToSlice[T any](rows *sqlx.Rows, result *[]interface{}) error {

	var err error
	*result, err = rows.SliceScan()
	if err != nil {
		return err
	}

	return nil
}

func FetchXRowsToMap[T any](rows *sqlx.Rows, result *map[string]interface{}) error {

	err := rows.MapScan(*result)
	if err != nil {
		return err
	}

	return nil
}
