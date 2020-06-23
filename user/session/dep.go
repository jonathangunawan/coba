package session

import (
	"time"

	"github.com/go-redis/redis"
)

type UserSessionItf interface {
	Store(SessionData) error
	Remove(string) error
}

type UserSession struct {
	conn *redis.Client
	exp  time.Duration
}

type SessionData struct {
	Key   string
	Token string
}

func NewUserSession(c *redis.Client) UserSession {
	return UserSession{
		conn: c,
		exp:  1200 * time.Second,
	}
}
