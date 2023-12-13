package db

import (
	env "aoc/server/config"
	"context"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
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

var db *sql.DB

func init() {
	prepareDatabase()
}

func prepareDatabase() {
	fmt.Println("prepare database...")
	env.LoadConfig()
	db = CreateConnection(env.Config.DBConnectionString)
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS requests (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(64),
    created_at TIMESTAMP,
    day INT,
    part INT,
    correct_answer BOOLEAN,
    valid BOOLEAN
);`)
	if err != nil {
		log.Fatal(err)
	}
}
