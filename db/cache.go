package database

import (
	"context"
	"log"

	redis "github.com/go-redis/redis/v8"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // If no password set
		DB:       0,                // Use default DB
	})

	res, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Connected to Redis:", res)
}
