package crud

import (
	"github.com/jmoiron/sqlx"
)

type UserCRUDItf interface {
	GetUser(...interface{}) (UserRow, error)
	InsertUser(UserRow) error
	UpdateUser(UserRow) error
	DeleteUser(int64) error
}

type UserCRUD struct {
	conn *sqlx.DB
}

type UserRow struct {
	UserID   int64  `db:"user_id"`
	Email    string `db:"email"`
	Address  string `db:"address"`
	Password string `db:"password"`
}

//init client to use User CRUD
func NewUserCRUD(c *sqlx.DB) UserCRUD {
	return UserCRUD{
		conn: c,
	}
}
