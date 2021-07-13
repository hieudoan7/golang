package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Go Redis Connection Test")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	fmt.Println(client)
}