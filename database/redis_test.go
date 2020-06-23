package database

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
)

func TestNewRedisCacheSuccess(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()

	dep, err := NewRedisCache(mr.Addr(), "")
	assert.Nil(t, err)
	assert.NotNil(t, dep)
}

func TestNewRedisCacheError(t *testing.T) {
	dep, err := NewRedisCache("", "")
	assert.NotNil(t, err)
	assert.Nil(t, dep)
}
