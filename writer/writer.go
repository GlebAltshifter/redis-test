package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	redis "github.com/go-redis/redis"
)

var client *redis.Client

func main() {
	// клиент БД Redis
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	cmd := redis.NewStringCmd("select", 0)
	err = client.Process(cmd)
	for {
		print("Input key: ")
		key, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		key = strings.TrimSpace(key)

		print("Input value: ")
		value, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		value = strings.TrimSpace(value)

		err = client.Set(key, value, 0).Err()
	}
}
