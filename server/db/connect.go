package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type Queryable interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func CreateConnection(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)
	if err != nil {
		panic(err)
	}
	return db
}
