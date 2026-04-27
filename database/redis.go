package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //local pake docker
		Password: "",               // kalau ada
		DB:       0,
	})

	return rdb
}
