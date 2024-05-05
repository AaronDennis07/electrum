package cache

import (
	"os"

	"github.com/go-redis/redis/v8"
)

type CacheClient struct {
	Redis *redis.Client
}

var Client CacheClient

func SetupCache() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	Client = CacheClient{
		Redis: redisClient,
	}
}
