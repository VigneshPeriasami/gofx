package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
)

type QueryClient struct {
	client DbClient
}

type QueryResult struct {
	db   *sql.DB
	rows *sql.Rows
}

func NewQuery(client DbClient) *QueryClient {
	return &QueryClient{
		client: client,
	}
}

func (q *QueryClient) Query(query string, args ...any) (*QueryResult, error) {
	db, err := q.client.Open()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return &QueryResult{
		db: db, rows: rows,
	}, err
}

var Module = fx.Options(
	fx.Provide(NewDbClient, NewQuery),
)
