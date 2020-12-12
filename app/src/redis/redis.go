package redis

import (
	"context"

	"github.com/go-redis/redis"
)

var ctx = context.Background()
var rdb *redis.Client

// Initialize connects to redis
func Initialize() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

// StoreToken stores csrf token with session id
func StoreToken(sessionID string, token string) (err error) {
	err = rdb.Set(ctx, sessionID, token, 0).Err()

	return
}

// GetToken gets csrf token from session id
func GetToken(sessionID string) (token string, err error) {
	token, err = rdb.Get(ctx, sessionID).Result()

	return
}
