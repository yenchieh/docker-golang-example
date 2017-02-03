package main

import (
	"fmt"

	redis "gopkg.in/redis.v5"
)

func main() {
	fmt.Println("-----Started--------")

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	err := client.Set("second", "2", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("first").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
