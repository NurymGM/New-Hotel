package initializers

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func ConnectToRedis() {
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	RDB = redis.NewClient(opts)

	_, err = RDB.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}