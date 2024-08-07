package database

import (
	"context"

	"github.com/awesomebfm/minepanel/pkg/auth"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	connStr string
	pool    *pgxpool.Pool
}

func NewDatabase(connStr string) (*Database, error) {
	pool, err := pgxpool.New(context.TODO(), connStr)
	if err != nil {
		return nil, err
	}

	return &Database{
		connStr: connStr,
		pool:    pool,
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

func (d *Database) PersistSession(session *auth.Session) error {
	query := `
		INSERT INTO sessions (user_id, hashed_secret, user_agent, ip_address, created_at, expires_at)
		VALUES (@userId, @hashedSecret, @userAgent, @ipAddress, @createdAt, @expiresAt)
		RETURNING id
		`
	args := pgx.NamedArgs{
		"userId":       session.UserId,
		"hashedSecret": session.HashedSecret,
		"userAgent":    session.UserAgent,
		"ipAddress":    session.IpAddress,
		"createdAt":    session.CreatedAt,
		"expiresAt":    session.ExpiresAt,
	}

	err := d.pool.QueryRow(context.TODO(), query, args).Scan(&session.Id)
	return err
}
