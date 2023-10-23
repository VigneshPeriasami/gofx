package database

func ReadRows[T any](q *QueryResult, mapperFn func(func(...any) error) T) []T {
	rows := q.rows
	result := []T{}
	for rows.Next() {
		result = append(result, mapperFn(rows.Scan))
	}
	defer q.db.Close()
	defer rows.Close()
	return result
}
