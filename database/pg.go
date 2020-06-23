package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// init PostgreDB
func NewPostgreDB(dbType, connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Open(dbType, connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
