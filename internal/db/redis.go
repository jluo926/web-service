package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func init() {
	log.Println("Connecting to redis")
	Redis = redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})
	str := Redis.Ping(context.Background())
	if str.Err() != nil {
		panic(str.Err())
	}
}
