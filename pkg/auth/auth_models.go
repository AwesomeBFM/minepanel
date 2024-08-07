package auth

import "time"

type User struct {
	Id             int
	Username       string
	HashedPassword string
	CreatedAt      time.Time
	LastLogin      time.Time
}

type Session struct {
	Id          int
	UserId      int
	HashedSecret string
	UserAgent   string
	IpAddress   string
	CreatedAt   time.Time
	ExpiresAt   time.Time
}
