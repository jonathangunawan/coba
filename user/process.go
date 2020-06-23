package user

import (
	"boyzgenk/coba/user/session"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

//Login Process using email & password
func (u User) Login(email, password string) error {
	res, err := u.crud.GetUser(email, password)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, userLog)
	}

	if err == sql.ErrNoRows {
		return errors.Wrap(UserNotFound, userLog)
	}

	//In this case I will store all of user data to redis
	//this action only for testing purpose
	//in production only not restricted data that can be store to redis
	ses := session.SessionData{
		Key:   fmt.Sprintf("%d", res.UserID),
		Token: fmt.Sprintf("%s%s%s", res.Email, res.Address, res.Password),
	}
	err = u.session.Store(ses)
	if err != nil {
		return errors.Wrap(err, userLog)
	}

	return nil
}

func (u User) Logout(id string) error {
	err := u.session.Remove(id)
	if err != nil {
		return errors.Wrap(err, userLog)
	}

	return nil
}
