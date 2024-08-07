package database

import (
	"context"

	"github.com/awesomebfm/minepanel/pkg/auth"
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

func (d *Database) FindUserByUsername(username string) (user *auth.User, err error) {
	row := d.pool.QueryRow(
		context.TODO(), 
		"SELECT id, hashed_password, created_at, last_login FROM users WHERE username = ?", 
		username)
	
	user.Username = username
	err = row.Scan(
		&user.Id,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.LastLogin,
	)

	if err != nil {
		return nil, err
	}
	return user, err
}

func (d *Database) FindSessionBySessionToken(token string)