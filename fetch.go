package godab

import "database/sql"

func FetchRowsTo[T any](rows *sql.Rows, result *[]T) error {
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

func FetchRowTo[T any](rows *sql.Row, result *T) error {
	err := rows.Scan(&result)
	if err != nil {
		return err
	}

	return nil
}
