package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgreDBSuccess(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	conn, err := NewPostgreDB("postgres", "testing")
	assert.Nil(t, err)
	assert.NotNil(t, conn)
}

func TestNewPostgreDBError(t *testing.T) {
	conn, err := NewPostgreDB("lala", "testing")
	assert.NotNil(t, err)
	assert.Nil(t, conn)
}
