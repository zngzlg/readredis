package readredis

import (
	"fmt"
	"context"
	"log"
	"sync"
	"runtime"
	"time"

	"github.com/go-redis/redis/v8"
)

func MyFunction() {
	fmt.Println("Hello, this is readredis module!")
}

// use context
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
	runtime.GOMAXPROCS(1)
	
	var wg sync.WaitGroup

	// Use goroutine to process each key
	for iter.Next(ctx) {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()
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

	// Wait for all goroutines to finish
	wg.Wait()
}

func process(value string) string {
	time.Sleep(time.Duration(1)*time.Second)
	return "processed " + value
}
