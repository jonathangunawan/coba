package user

import (
	"boyzgenk/coba/user/crud"
	"boyzgenk/coba/user/session"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPG struct{ mock.Mock }

func (m *MockPG) GetUser(a ...interface{}) (crud.UserRow, error) {
	args := m.Called(a)
	return args.Get(0).(crud.UserRow), args.Error(1)
}

func (m *MockPG) InsertUser(a crud.UserRow) error {
	args := m.Called(a)
	return args.Error(0)
}

func (m *MockPG) UpdateUser(a crud.UserRow) error {
	args := m.Called(a)
	return args.Error(0)
}

func (m *MockPG) DeleteUser(a int64) error {
	args := m.Called(a)
	return args.Error(0)
}

type MockRds struct{ mock.Mock }

func (m *MockRds) Store(a session.SessionData) error {
	args := m.Called(a)
	return args.Error(0)
}

func (m *MockRds) Remove(a string) error {
	args := m.Called(a)
	return args.Error(0)
}

func TestUser_LoginSuccess(t *testing.T) {
	mockPG := new(MockPG)
	mockRds := new(MockRds)

	pgData := crud.UserRow{
		UserID:   1,
		Email:    "lala@mail.com",
		Address:  "disana",
		Password: "1234",
	}

	rdsData := session.SessionData{
		Key:   "1",
		Token: "lala@mail.comdisana1234",
	}

	mockPG.On("GetUser", []interface{}{"lala@mail.com", "1234"}).Return(pgData, nil)
	mockRds.On("Store", rdsData).Return(nil)

	dep := User{
		crud:    mockPG,
		session: mockRds,
	}

	err := dep.Login("lala@mail.com", "1234")
	assert.Nil(t, err)
}

func TestUser_LoginErrorStore(t *testing.T) {
	mockPG := new(MockPG)
	mockRds := new(MockRds)

	pgData := crud.UserRow{
		UserID:   1,
		Email:    "lala@mail.com",
		Address:  "disana",
		Password: "1234",
	}

	rdsData := session.SessionData{
		Key:   "1",
		Token: "lala@mail.comdisana1234",
	}

	mockPG.On("GetUser", []interface{}{"lala@mail.com", "1234"}).Return(pgData, nil)
	mockRds.On("Store", rdsData).Return(fmt.Errorf("error"))

	dep := User{
		crud:    mockPG,
		session: mockRds,
	}

	err := dep.Login("lala@mail.com", "1234")
	assert.NotNil(t, err)
}

func TestUser_LoginErrorNoRows(t *testing.T) {
	mockPG := new(MockPG)
	mockRds := new(MockRds)

	pgData := crud.UserRow{
		UserID:   1,
		Email:    "lala@mail.com",
		Address:  "disana",
		Password: "1234",
	}

	mockPG.On("GetUser", []interface{}{"lala@mail.com", "1234"}).Return(pgData, sql.ErrNoRows)

	dep := User{
		crud:    mockPG,
		session: mockRds,
	}

	err := dep.Login("lala@mail.com", "1234")
	assert.NotNil(t, err)
}

func TestUser_LoginErrorGetUser(t *testing.T) {
	mockPG := new(MockPG)
	mockRds := new(MockRds)

	pgData := crud.UserRow{
		UserID:   1,
		Email:    "lala@mail.com",
		Address:  "disana",
		Password: "1234",
	}

	mockPG.On("GetUser", []interface{}{"lala@mail.com", "1234"}).Return(pgData, fmt.Errorf("error"))

	dep := User{
		crud:    mockPG,
		session: mockRds,
	}

	err := dep.Login("lala@mail.com", "1234")
	assert.NotNil(t, err)
}

func TestUser_LogoutSuccess(t *testing.T) {
	mockRds := new(MockRds)

	mockRds.On("Remove", "1").Return(nil)

	dep := User{
		session: mockRds,
	}

	err := dep.Logout("1")
	assert.Nil(t, err)
}

func TestUser_LogoutErrorRemove(t *testing.T) {
	mockRds := new(MockRds)

	mockRds.On("Remove", "1").Return(fmt.Errorf("error"))

	dep := User{
		session: mockRds,
	}

	err := dep.Logout("1")
	assert.NotNil(t, err)
}
