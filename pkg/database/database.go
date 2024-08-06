package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	connStr string
	pool *pgxpool.Pool
}

func NewDatabase(connStr string) (*Database, error) {
	pool, err := pgxpool.New(context.TODO(), connStr)
	if err != nil {
		return nil, err
	}

	return &Database{
		connStr: connStr,
		pool: pool,
	}, nil
}

func (d *Database) Close() {
	// Additional closing logic
	d.pool.Close()
}