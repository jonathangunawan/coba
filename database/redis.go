package database

import "github.com/go-redis/redis"

func NewRedisCache(addr, password string) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	_, err := conn.Ping().Result()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
