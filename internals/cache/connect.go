package cache

import (
	"context"
	"log"
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
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis")
	}

	Client = CacheClient{
		Redis: redisClient,
	}
}
