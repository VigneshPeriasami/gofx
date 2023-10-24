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

type queryBuilder func(values int) string

func insertQuery(tableName string, columnArr []string) queryBuilder {
	columns := strings.Join(columnArr, ",")
	return func(rows int) string {
		return fmt.Sprintf(
			"INSERT into %s (%s) VALUES %s", tableName,
			columns, queryPlaces(rows, len(columnArr)),
		)
	}
}

type MapperFn[T any] func(T) []interface{}

type InsertQ[T any] struct {
	TableName string
	Columns   []string
	MapperFn  MapperFn[T]
}

func (q *InsertQ[T]) Build(data []T) (string, []interface{}) {
	return insertQuery(q.TableName, q.Columns)(len(data)), Flatten[T](data, q.MapperFn)
}
