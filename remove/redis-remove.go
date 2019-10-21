package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := newClient()

	err := client.FlushAll().Err()
	if err != nil {
		fmt.Println(err)
	}
}

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
