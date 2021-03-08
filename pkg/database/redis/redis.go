package redis

import (
	"github.com/go-redis/redis/v8"
)

func ConnectRedis(addr, password string, db int) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return rdb
}
