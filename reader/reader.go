package main

import (
	"log"

	"github.com/go-redis/redis"
)

var client *redis.Client

type data struct {
	Key string
	Val string
}

func main() {

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	// выбираем БД №0 - разложенные числа
	cmd := redis.NewStringCmd("select", 0)
	err := client.Process(cmd)
	if err != nil {
		log.Fatal(err)
	}

	// получаем все ключи из БД
	keys := client.Keys("*")

	// var solved []data
	var item data

	// для каждого ключа получаем значение и добавляем в массив
	for _, key := range keys.Val() {
		item.Key = key
		val, err := client.Get(key).Result()
		if err != nil {
			log.Fatal(err)
		}
		item.Val = val

		log.Print(item)

		// solved = append(solved, item)

	}
}
