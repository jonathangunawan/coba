package user

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"
)

func TestNewUser(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	dep := NewUser(sqlxDB, nil)
	assert.NotNil(t, dep)
}
