package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // adjust for your environment
		Password: "",               // no password set
		DB:       0,                // default DB
	})
}
