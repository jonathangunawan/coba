package user

import (
	"boyzgenk/coba/user/crud"
	"boyzgenk/coba/user/session"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type UserItf interface {
	Login(string, string) error
	Logout(string) error
}

type User struct {
	crud    crud.UserCRUDItf
	session session.UserSessionItf
}

func NewUser(c *sqlx.DB, s *redis.Client) User {
	return User{
		crud:    crud.NewUserCRUD(c),
		session: session.NewUserSession(s),
	}
}
