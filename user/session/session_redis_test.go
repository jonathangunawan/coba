package session

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

func TestUserSession_StoreSuccess(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()

	c := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	data := SessionData{
		Key:   "1",
		Token: "lala",
	}
	dep := NewUserSession(c)
	err = dep.Store(data)
	assert.Nil(t, err)
	mr.CheckGet(t, "1", "lala")
}

func TestUserSession_StoreError(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()

	c := redis.NewClient(&redis.Options{
		Addr: "",
	})

	data := SessionData{
		Key:   "1",
		Token: "lala",
	}
	dep := NewUserSession(c)
	err = dep.Store(data)
	assert.NotNil(t, err)
}

func TestUserSession_RemoveSuccess(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()

	c := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mr.Set("1", "lala")
	dep := NewUserSession(c)
	err = dep.Remove("1")
	assert.Nil(t, err)
	_, getErr := mr.Get("1")
	assert.Equal(t, "ERR no such key", getErr.Error())
}

func TestUserSession_RemoveError(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()

	c := redis.NewClient(&redis.Options{
		Addr: "",
	})

	mr.Set("1", "lala")
	dep := NewUserSession(c)
	err = dep.Remove("1")
	assert.NotNil(t, err)
	_, getErr := mr.Get("1")
	assert.Nil(t, getErr)
}
