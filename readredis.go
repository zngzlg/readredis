package readredis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func MyFunction() {
	fmt.Println("Hello, this is readredis module!")
}

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
	key := "example_key"
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Error getting value from Redis: %v", err)
	}

	// Process the data
	result := process(value)

	// Return the result
	fmt.Printf("Processed result: %s\n", result)
}

func process(value string) string {
	return "processed"
}
