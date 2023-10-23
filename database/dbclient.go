package database

import (
	"database/sql"
	"time"

	"github.com/vigneshperiasami/analytics/environment"
)

type DbClient struct {
	conn string
}

func NewDbClient(config *environment.ConfigResult) *DbClient {
	return &DbClient{
		conn: config.DbConn,
	}
}

func (dbClient *DbClient) Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbClient.conn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, err
}
