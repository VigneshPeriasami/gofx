package database

import (
	"database/sql"
	"time"

	"go.uber.org/fx"
)

type DbClient interface {
	Open() (*sql.DB, error)
}

type DbClientResult struct {
	conn string
}

type DbParams struct {
	fx.In
	Conn string `name:"dbconn"`
}

func NewDbClient(dbParams DbParams) DbClient {
	return &DbClientResult{
		conn: dbParams.Conn,
	}
}

func (dbClient *DbClientResult) Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbClient.conn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, err
}
