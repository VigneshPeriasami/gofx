package sqlhelper

import (
	"fmt"
	"strings"
)

func Times[T any](repeat T, times int) []T {
	result := []T{}
	for i := 0; i < times; i++ {
		result = append(result, repeat)
	}
	return result
}

func queryPlaces(rows, columns int) string {
	columnPlaces := Times[string]("?", columns)

	columnStr := fmt.Sprintf("(%s)", strings.Join(columnPlaces, ","))
	queryPlaces := Times[string](columnStr, rows)

	return strings.Join(queryPlaces, ",")
}

func Flatten[T any, R any](modelList []T, fn func(T) []R) []R {
	flatten := []R{}
	for _, v := range modelList {
		flatten = append(flatten, fn(v)...)
	}
	return flatten
}

type QueryBuilder func(values int) string

func InsertQuery(tableName string, columnArr []string) QueryBuilder {
	columns := strings.Join(columnArr, ",")
	return func(rows int) string {
		return fmt.Sprintf(
			"INSERT into %s (%s) VALUES %s", tableName,
			columns, queryPlaces(rows, len(columnArr)),
		)
	}
}
