package readredis

import (
	"fmt"
	"context"
	"log"
	"github.com/go-redis/redis/v8"
)

func MyFunction() {
	fmt.Println("Hello, this is readredis module!")
}
var ctx = context.Background()

func ProcessData() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Check the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	// Read data from Redis
	iter := rdb.Scan(ctx,0,"*",0).Iterator()

	// Use goroutine to process each key
	for iter.Next(ctx) {
		go func(key string) {
			value, err := rdb.Get(ctx, key).Result()
			if err != nil {
				log.Fatalf("Error getting value from Redis: %v", err)
			}
			result := process(value)
			fmt.Printf("Processed result: %s\n", result)
		}(iter.Val())
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating keys: %v", err)
	}
}

func process(value string) string {
	return "processed " + value
}
