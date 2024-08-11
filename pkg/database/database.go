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

func (d *Database) FindUserByUsername(username string) (*auth.User, error) {
	row := d.pool.QueryRow(
		context.TODO(),
		"SELECT id, hashed_password, created_at, last_login FROM users WHERE username = $1",
		username)

	var user auth.User
	user.Username = username
	err := row.Scan(
		&user.Id,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.LastLogin,
	)

	return &user, err
}

func (d *Database) PersistSession(session *auth.Session) error {
	query := `
		INSERT INTO sessions (user_id, hashed_secret, user_agent, ip_address, created_at, expires_at)
		VALUES (@user_id, @hashed_secret, @user_agent, @ip_address, @created_at, @expires_at)
		RETURNING id
		`
	args := pgx.NamedArgs{
		"user_id":       session.UserId,
		"hashed_secret": session.HashedSecret,
		"user_agent":    session.UserAgent,
		"ip_address":    session.IpAddress,
		"created_at":    session.CreatedAt,
		"expires_at":    session.ExpiresAt,
	}

	return d.pool.QueryRow(context.TODO(), query, args).Scan(&session.Id)
}

func (d *Database) FindSessionById(id int) (*auth.Session, error) {
	row := d.pool.QueryRow(context.TODO(), "SELECT * FROM sessions WHERE id = $1", id)

	var session auth.Session
	err := row.Scan(
		&session.Id,
		&session.UserId,
		&session.HashedSecret,
		&session.UserAgent,
		&session.IpAddress,
		&session.CreatedAt,
		&session.ExpiresAt,
	)

	return &session, err
}

func (d *Database) PersistUser(user *auth.User) error {
	query := `
		INSERT INTO users (username, hashed_password, created_at, last_login)
		VALUES (@username, @hashed_password, @created_at, @last_login)
		RETURNING id
		`
	args := pgx.NamedArgs{
		"username":        user.Username,
		"hashed_password": user.HashedPassword,
		"created_at":      user.CreatedAt,
		"last_login":      user.LastLogin,
	}

	err := d.pool.QueryRow(context.TODO(), query, args).Scan(&user.Id)

	return err
}
