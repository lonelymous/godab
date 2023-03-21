package godab

import "database/sql"

func FetchRowsTo(rows sql.Rows, a *[]any) (error, []any) {
	var result []any

	for rows.Next() {
		var row any

		err := rows.Scan(&row)
		if err != nil {
			return err, nil
		}

		result = append(result, row)
	}

	return nil, result
}
