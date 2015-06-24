package main

import (
	"gopkg.in/redis.v3"
	"log"
	"os"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: getRedisAddress(),
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalln("Error connecting to redis", err)
	} else {
		log.Println("Connected to redis")
	}
}

func getRedisAddress() (cs string) {
	cs = "localhost:6379"
	if envvar := os.Getenv("REDIS_ADDR"); envvar != "" {
		cs = envvar
	}
	return
}
