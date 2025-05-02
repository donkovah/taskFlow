package messaging

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),     // e.g., "localhost:6379"
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func Publish(channel string, message string) {
	err := rdb.Publish(ctx, channel, message).Err()
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}
}

func Subscribe(channel string) {
	pubsub := rdb.Subscribe(ctx, channel)

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// Consume messages.
	go func() {
		for msg := range ch {
			log.Printf("Received a message: %s", msg.Payload)
			// Process the message
		}
	}()
}
