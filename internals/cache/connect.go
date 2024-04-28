package cache

import "github.com/go-redis/redis/v8"

type CacheClient struct {
	Redis *redis.Client
}

var Client CacheClient

func SetupCache() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Client = CacheClient{
		Redis: redisClient,
	}
}
