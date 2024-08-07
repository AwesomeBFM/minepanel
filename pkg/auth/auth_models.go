package auth

import "time"

type User struct {
	Id             int64
	Username       string
	HashedPassword string
	CreatedAt      time.Time
	LastLogin      time.Time
}

type Session struct {
	Id          int64
	UserId      int64
	HashedToken string
	UserAgent   string
	IpAddress   string
	CreatedAt   time.Time
	ExpiresAt   time.Time
}
