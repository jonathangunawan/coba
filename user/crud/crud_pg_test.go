package crud

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserCRUD_GetUserSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{
		"user_id", "email", "address", "password",
	}).AddRow(
		1, "lala@mail.com", "disana", "1234",
	)
	mock.ExpectQuery("SELECT (.+)").WithArgs("lala@mail.com", "1234").WillReturnRows(rows)

	dep := NewUserCRUD(sqlxDB)
	res, err := dep.GetUser("lala@mail.com", "1234")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), res.UserID)
}

func TestUserCRUD_GetUserError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectQuery("SELECT (.+)").WithArgs("lala@mail.com", "1234").WillReturnError(fmt.Errorf("error"))

	dep := NewUserCRUD(sqlxDB)
	res, err := dep.GetUser("lala@mail.com", "1234")
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), res.UserID)
}

func TestUserCRUD_InsertUserSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectExec("INSERT (.+)").WithArgs("lala@mail.com", "disana", "123").WillReturnResult(sqlmock.NewResult(1, 1))

	dep := NewUserCRUD(sqlxDB)
	data := UserRow{
		Email:    "lala@mail.com",
		Address:  "disana",
		Password: "123",
	}
	err := dep.InsertUser(data)
	assert.Nil(t, err)
}

func TestUserCRUD_InsertUserError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectExec("INSERT (.+)").WithArgs("lala@mail.com", "disana", "123").WillReturnError(fmt.Errorf("error"))

	dep := NewUserCRUD(sqlxDB)
	data := UserRow{
		Email:    "lala@mail.com",
		Address:  "disana",
		Password: "123",
	}
	err := dep.InsertUser(data)
	assert.NotNil(t, err)
}

func TestUserCRUD_UpdateUserSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectExec("UPDATE (.+)").WithArgs("lala@mail.com", "disana", "123", int64(1)).WillReturnResult(sqlmock.NewResult(1, 1))

	dep := NewUserCRUD(sqlxDB)
	data := UserRow{
		UserID:   1,
		Email:    "lala@mail.com",
		Address:  "disana",
		Password: "123",
	}
	err := dep.UpdateUser(data)
	assert.Nil(t, err)
}

func TestUserCRUD_UpdateUserError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectExec("UPDATE (.+)").WithArgs("lala@mail.com", "disana", "123", int64(1)).WillReturnError(fmt.Errorf("error"))

	dep := NewUserCRUD(sqlxDB)
	data := UserRow{
		UserID:   1,
		Email:    "lala@mail.com",
		Address:  "disana",
		Password: "123",
	}
	err := dep.UpdateUser(data)
	assert.NotNil(t, err)
}

func TestUserCRUD_DeleteUserSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectExec("DELETE (.+)").WithArgs(int64(1)).WillReturnResult(sqlmock.NewResult(1, 1))

	dep := NewUserCRUD(sqlxDB)
	err := dep.DeleteUser(1)
	assert.Nil(t, err)
}

func TestUserCRUD_DeleteUserError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectExec("DELETE (.+)").WithArgs(int64(1)).WillReturnError(fmt.Errorf("error"))

	dep := NewUserCRUD(sqlxDB)
	err := dep.DeleteUser(1)
	assert.NotNil(t, err)
}

func TestUserCRUD_CreateTableSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectExec("CREATE (.+)").WillReturnResult(sqlmock.NewResult(1, 1))

	dep := NewUserCRUD(sqlxDB)
	err := dep.CreateTable()
	assert.Nil(t, err)
}

func TestUserCRUD_CreateTableError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectExec("CREATE (.+)").WillReturnError(fmt.Errorf("error"))

	dep := NewUserCRUD(sqlxDB)
	err := dep.CreateTable()
	assert.NotNil(t, err)
}
